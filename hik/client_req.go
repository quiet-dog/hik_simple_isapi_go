package hk_gateway

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	Post   = "POST"
	Get    = "GET"
	Put    = "PUT"
	Delete = "DELETE"
)

type ReqInitParam struct {
	Url    string
	Query  map[string]string
	Body   interface{}
	Result interface{}
	Method string
}

func (c *HikClient) Do(param ReqInitParam) (err error) {

	if !c.isConnect && param.Url != "/ISAPI/Security/userCheck" {
		return
	}

	req := c.client.R().SetDigestAuth(c.hikConfig.Username, c.hikConfig.Password).SetDebug(true)

	req = req.SetQueryParams(param.Query).SetBody(param.Body)

	var resp *resty.Response
	switch param.Method {
	case Post:
		{
			resp, err = req.Post(param.Url)
		}
	case Get:
		{
			resp, err = req.Get(param.Url)
		}
	case Put:
		{
			resp, err = req.Put(param.Url)
		}
	case Delete:
		{
			resp, err = req.Delete(param.Url)
		}
	}

	if err != nil {
		c.isConnect = false
		return
	}

	// 判断响应的内容格式
	switch resp.Header().Get("Content-Type") {
	case "application/xml":
		{
			if resp.StatusCode() != 200 {
				userCheck := UserCheckRes{}
				// 先验证是否经过摘要认证
				if err = xml.Unmarshal(resp.Body(), &userCheck); err != nil {
					fmt.Printf("xml 解析失败 %s", err.Error())
					return fmt.Errorf("xml 解析失败 %s", err.Error())
				}

				if userCheck.StatusValue != 0 && userCheck.StatusValue != 200 {
					err = fmt.Errorf("StatusString = %s, StatusValue = %d, RetryLoginTime = %d", userCheck.StatusString, userCheck.StatusValue, userCheck.RetryLoginTime)
					c.isConnect = false
					fmt.Println(err.Error())
					return
				}

			}

			c.isConnect = true
			if err = xml.Unmarshal(resp.Body(), param.Result); err != nil {
				fmt.Printf("xml 解析失败 验证后  %s", err.Error())
				return fmt.Errorf("xml 解析失败 验证后1  %s", err.Error())
			}
		}
	case "application/json":
		{
			// if resp.StatusCode() != 200 {
			errMsg := ErrorMsg{}
			if err = json.Unmarshal(resp.Body(), &errMsg); err != nil {
				return errors.New("json 解析失败" + err.Error())
			}

			// 判断返回的json格式是不是这个格式
			if errMsg.StatusCode != 0 && errMsg.StatusCode != 1 && errMsg.SubStatusCode != "OK" {
				err = fmt.Errorf("ErrorCode = %d, ErrorMsg = %s, StatusString = %s, StatusCode = %d, SubStatusCode = %s", errMsg.ErrorCode, errMsg.ErrorMsg, errMsg.StatusString, errMsg.StatusCode, errMsg.SubStatusCode)
				return
			}

			if err = json.Unmarshal(resp.Body(), param.Result); err != nil {
				return errors.New("json 解析失败 验证后2" + err.Error() + resp.Status())
			}
		}
	case "image/jpeg":
		{
			c.isConnect = true
			param.Result.(*bytes.Buffer).Write(resp.Body())
		}
	}

	return
}
