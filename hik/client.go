package hik

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/go-resty/resty/v2"
)

type HikConfig struct {
	Ip       string
	Port     int
	Username string
	Password string
}

type HikClient struct {
	hikConfig HikConfig
	client    *resty.Client
	isConnect bool
	// 上下文
	ctx context.Context
	// 取消上下文
	cancel context.CancelFunc
	// 长连接状态
	longConnect bool
}

func newHikClinet(conf HikConfig) (hikClient *HikClient, err error) {
	hikClient = &HikClient{
		hikConfig: conf,
		client:    resty.New().SetBaseURL(fmt.Sprintf("http://%s:%d", conf.Ip, conf.Port)).SetTimeout(3 * time.Second),
		isConnect: false,
	}

	_, err = hikClient.UserCheck()
	// 开启长连接

	return hikClient, err
}

func (c *HikClient) UserCheck() (result *UserCheckRes, err error) {
	req := ReqInitParam{
		Url:    "/ISAPI/Security/userCheck",
		Query:  nil,
		Body:   nil,
		Result: &result,
		Method: Get,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

func (c *HikClient) AddPersonInfo(person AddPersonInfoReq) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/UserInfo/Record",
		Query: map[string]string{
			"format": "json",
		},
		Body:   person,
		Result: &result,
		Method: Post,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 设置人脸信息
// /ISAPI/Intelligent/FDLib/FDSetUp
func (c *HikClient) SetFaceInfo(reqBody SetFaceInfo) (result *ErrorMsg, err error) {

	//  将reqBody转换成多表单 multipart/form-data
	var bodyBuf bytes.Buffer
	writer := multipart.NewWriter(&bodyBuf)
	defer writer.Close()

	jsonStr, err := json.Marshal(reqBody.FaceDataRecord)
	if err != nil {
		return
	}
	if err := writer.WriteField("FaceDataRecord", string(jsonStr)); err != nil {
		return nil, err
	}
	buf, err := writer.CreateFormFile("img", reqBody.Img.Filename)
	if err != nil {
		return nil, err
	}

	img, err := reqBody.Img.Open()
	if err != nil {
		return nil, err
	}
	defer img.Close()
	if _, err = io.Copy(buf, img); err != nil {
		return nil, err
	}

	req := ReqInitParam{
		Url: "/ISAPI/Intelligent/FDLib/FDSetUp",
		Query: map[string]string{
			"format": "json",
		},
		Body:   bodyBuf,
		Result: &result,
		Method: Post,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 修改人员信息
// ISAPI/AccessControl/UserInfo/Modify?format=json
func (c *HikClient) ModifyPersonInfo(person AddPersonInfoReq) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/UserInfo/Modify",
		Query: map[string]string{
			"format": "json",
		},
		Body:   person,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 删除人员信息
// ISAPI/AccessControl/UserInfo/Delete?format=json
func (c *HikClient) DeletePersonInfo(reqBody DelPersonInfo) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/UserInfo/Delete",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 获取人员列表
// /ISAPI/AccessControl/UserInfo/Search
func (c *HikClient) GetPersonInfoList(reqBody GetPersonInfoList) (result *UserInfoSearch, err error) {

	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/UserInfo/Search",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Post,
	}

	if err := c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return
}

// 日志查询
// /ISAPI/ContentMgmt/logSearch
func (c *HikClient) LogSearch(reqBody LogSearch) (result *LogSearchRes, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/ContentMgmt/logSearch",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Post,
	}

	if err := c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return result, nil
}

// 获取人员数量信息
// /ISAPI/AccessControl/UserInfo/Count
func (c *HikClient) GetPersonInfoCount() (result *GetPersonInfoCountRes, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/UserInfo/Count",
		Query: map[string]string{
			"format": "json",
		},
		Body:   nil,
		Result: &result,
		Method: Get,
	}

	if err := c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return
}

// 查询设备中已有的人脸数量及人脸信息
// /ISAPI/Intelligent/FDLib/Count?format=json
func (c *HikClient) GetFaceInfoCount() (result *GetFaceInfoCountRes, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/Intelligent/FDLib/Count",
		Query: map[string]string{
			"format": "json",
		},
		Body:   nil,
		Result: &result,
		Method: Get,
	}

	if err := c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return result, nil
}

// 查询指定或全部人员的卡数量
// /ISAPI/AccessControl/CardInfo/Count?format=json
func (c *HikClient) GetCardInfoCount() (result *GetCardInfoCountRes, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/CardInfo/Count",
		Query: map[string]string{
			"format": "json",
		},
		Body:   nil,
		Result: &result,
		Method: Get,
	}

	if err := c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return
}

// 查询门禁事件条数。
// /ISAPI/AccessControl/AcsEventTotalNum?format=json
func (c *HikClient) GetAcsEventTotalNum(reqBody GetAcsEventTotalNumReq) (result *GetAcsEventTotalNumRes, err error) {
	body := make(map[string]interface{})
	body["AcsEventTotalNumCond"] = make(map[string]interface{})
	subBody := body["AcsEventTotalNumCond"].(map[string]interface{})
	subBody["major"] = reqBody.AcsEventTotalNumCond.Major
	subBody["minor"] = reqBody.AcsEventTotalNumCond.Minor

	if reqBody.AcsEventTotalNumCond.CardNo != "" {
		subBody["cardNo"] = reqBody.AcsEventTotalNumCond.CardNo
	}
	if reqBody.AcsEventTotalNumCond.Name != "" {
		subBody["name"] = reqBody.AcsEventTotalNumCond.Name
	}
	if reqBody.AcsEventTotalNumCond.EmployeeNoString != "" {
		subBody["employeeNoString"] = reqBody.AcsEventTotalNumCond.EmployeeNoString
	}
	if reqBody.AcsEventTotalNumCond.BeginSerialNo != 0 {
		subBody["beginSerialNo"] = reqBody.AcsEventTotalNumCond.BeginSerialNo
	}
	if reqBody.AcsEventTotalNumCond.EndSerialNo != 0 {
		subBody["endSerialNo"] = reqBody.AcsEventTotalNumCond.EndSerialNo
	}

	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/AcsEventTotalNum",
		Query: map[string]string{
			"format": "json",
		},
		Body:   body,
		Result: &result,
		Method: Post,
	}

	if err := c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return
}

// 获取门禁主机工作状态
// /ISAPI/AccessControl/AcsWorkStatus?format=json
func (c *HikClient) GetAcsWorkStatus() (result *GetAcsWorkStatusRes, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/AcsWorkStatus",
		Query: map[string]string{
			"format": "json",
		},
		Body:   nil,
		Result: &result,
		Method: Get,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return nil, err
	}
	c.isConnect = true
	return result, nil
}

// 添加卡信息
// /ISAPI/AccessControl/CardInfo/Record
func (c *HikClient) AddCardInfo(reqBody AddCardInfoReq) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/CardInfo/Record",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Post,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 获取用户的图片
func (c *HikClient) GetPicture(reqBody GetPictureReq) (result []byte, err error) {
	// 创建新的缓冲区
	resultBuf := bytes.NewBuffer([]byte{})
	req := ReqInitParam{
		Url:    reqBody.URL,
		Query:  nil,
		Body:   nil,
		Result: resultBuf,
		Method: Post,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}

	result = make([]byte, resultBuf.Len())
	if _, err = resultBuf.Read(result); err != nil {
		return
	}
	c.isConnect = true
	return result, nil
}

// 删除卡信息
// /ISAPI/AccessControl/CardInfo/Delete?format=json
func (c *HikClient) DeleteCardInfo(reqBody DelCardInfoReq) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/CardInfo/Delete",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 获取卡信息
// /ISAPI/AccessControl/CardInfo/Search?format=json
func (c *HikClient) GetCardInfo(reqBody GetCardInfoReq) (result *GetCardInfoRes, err error) {

	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/CardInfo/Search",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Post,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 查询门禁事件
// /ISAPI/AccessControl/AcsEvent?format=json
func (c *HikClient) GetAcsEvent(reqBody GetAcsEventReq) (result *GetAcsEventRes, err error) {

	body := make(map[string]interface{})
	body["AcsEventCond"] = map[string]interface{}{
		"searchID":             reqBody.AcsEventCond.SearchID,
		"searchResultPosition": reqBody.AcsEventCond.SearchResultPosition,
		"maxResults":           reqBody.AcsEventCond.MaxResults,
		"major":                reqBody.AcsEventCond.Major,
		"minor":                reqBody.AcsEventCond.Minor,
		"timeReverseOrder":     true,
	}

	if reqBody.AcsEventCond.StartTime != "" {
		body["AcsEventCond"].(map[string]interface{})["startTime"] = reqBody.AcsEventCond.StartTime
	}

	if reqBody.AcsEventCond.EndTime != "" {
		body["AcsEventCond"].(map[string]interface{})["endTime"] = reqBody.AcsEventCond.EndTime
	}
	if reqBody.AcsEventCond.CardNo != "" {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["cardNo"] = reqBody.AcsEventCond.CardNo
	}
	if reqBody.AcsEventCond.Name != "" {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["name"] = reqBody.AcsEventCond.Name
	}
	if reqBody.AcsEventCond.EmployeeNoString != "" {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["employeeNoString"] = reqBody.AcsEventCond.EmployeeNoString
	}

	if reqBody.AcsEventCond.BeginSerialNo != 0 {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["beginSerialNo"] = reqBody.AcsEventCond.BeginSerialNo
	}

	if reqBody.AcsEventCond.EndSerialNo != 0 {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["endSerialNo"] = reqBody.AcsEventCond.EndSerialNo
	}

	if reqBody.AcsEventCond.Major != 0 {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["major"] = reqBody.AcsEventCond.Major
	}

	if reqBody.AcsEventCond.Minor != 0 {
		// 删除空字段
		body["AcsEventCond"].(map[string]interface{})["minor"] = reqBody.AcsEventCond.Minor
	}

	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/AcsEvent",
		Query: map[string]string{
			"format": "json",
		},
		Body:   body,
		Result: &result,
		Method: Post,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 配置门禁主机参数
// /ISAPI/AccessControl/AcsCfg?format=json
func (c *HikClient) SetAcsCfg(reqBody SetAcsCfgReq) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/AcsCfg",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// /ISAPI/AccessControl/AcsEvent/StorageCfg?format=json
func (c *HikClient) SetStorageCfg(reqBody SetStorageCfgReq) (result *ErrorMsg, err error) {
	body := make(map[string]interface{})
	body["EventStorageCfg"] = make(map[string]interface{})
	subBody := body["EventStorageCfg"].(map[string]interface{})
	if reqBody.EventStorageCfg.CheckTime != "" {
		subBody["checkTime"] = reqBody.EventStorageCfg.CheckTime
	}
	if reqBody.EventStorageCfg.Mode != "" {
		subBody["mode"] = reqBody.EventStorageCfg.Mode
	}
	if reqBody.EventStorageCfg.Period != 0 {
		subBody["period"] = reqBody.EventStorageCfg.Period
	}

	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/AcsEvent/StorageCfg",
		Query: map[string]string{
			"format": "json",
		},
		Body:   body,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 配置人员及凭证显示参数
// /ISAPI/AccessControl/userAndRightShow?format=json
func (c *HikClient) GetUserAndRightShow(reqBody GetUserAndRightShowReq) (result *ErrorMsg, err error) {
	body := make(map[string]interface{})
	body["showAuthenticationList"] = reqBody.ShowAuthenticationList
	if reqBody.ShowCardNo != "" {
		body["showCardNo"] = reqBody.ShowCardNo
	}
	if reqBody.ShowDuration != 0 {
		body["showDuration"] = reqBody.ShowDuration
	}

	req := ReqInitParam{
		Url: "/ISAPI/AccessControl/userAndRightShow",
		Query: map[string]string{
			"format": "json",
		},
		Body:   reqBody,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 远程控门
// PUT/ISAPI/AccessControl/RemoteControl/door/<doorID>
func (c *HikClient) RemoteControlDoor(reqBody RemoteControlDoorReq) (result *ErrorMsg, err error) {
	req := ReqInitParam{
		Url:    fmt.Sprintf("/ISAPI/AccessControl/RemoteControl/door/%s", reqBody.DoorID),
		Query:  map[string]string{},
		Body:   reqBody.RemoteControlDoor,
		Result: &result,
		Method: Put,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return
}

// 获取设备信息参数
// GET/ISAPI/System/deviceInfo
func (c *HikClient) GetDeviceInfo() (result *DeviceInfo, err error) {
	req := ReqInitParam{
		Url:    "/ISAPI/System/deviceInfo",
		Query:  nil,
		Body:   nil,
		Result: &result,
		Method: Get,
	}

	if err = c.Do(req); err != nil {
		c.isConnect = false
		return
	}
	c.isConnect = true
	return result, nil
}
