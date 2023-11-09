package hik

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type HikGateway struct {
	hikMap sync.Map
	// 注册接受广播的通道
	broadClient sync.Map
}

type HikInfo struct {
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsConnect   bool   `json:"is_connect"`
	LongConnect bool   `json:"long_connect"`
}

func NewHikGateway() *HikGateway {
	return &HikGateway{}
}

// 注册一个Hik服务
func (h *HikGateway) RegisterHikGateway(hikConfig HikConfig) (err error) {

	// 判断是否已经注册
	if _, ok := h.hikMap.Load(hikConfig.Ip); ok {
		return errors.New("已经注册")
	}

	hikClient, err := newHikClinet(hikConfig)
	if err != nil {
		return
	}

	h.hikMap.Store(hikConfig.Ip, hikClient)

	// 接收广播
	hikClient.ctx, hikClient.cancel = context.WithCancel(context.Background())
	go h.handelBoardcast(hikClient)

	// h.longMap.Store(hikConfig.Ip, cancel)

	// // 注册任务服务
	// task := gocron.NewScheduler()
	// task.Every(20).Seconds().Do(h.UserCheck, hikConfig)
	// <-task.Start()
	// h.taskMap.Store(hikConfig, task)
	return
}

// 删除一个Hik服务
func (h *HikGateway) DeleteHikGateway(key string) {

	// 关闭长连接
	// h.cancelBoardcast(key)
	// 删除服务
	h.cancelHikGateway(key)
}

func (h *HikGateway) cancelHikGateway(key string) {
	if v, ok := h.hikMap.Load(key); ok {
		// 关闭长连接
		v.(*HikClient).cancel()
		h.hikMap.Delete(key)
	}
}

// 更新一个Hik服务
func (h *HikGateway) UpdateHikGateway(hikConfig HikConfig) (err error) {
	// 获取原来的服务
	hikConfig.Ip = strings.Replace(hikConfig.Ip, "http://", "", -1)
	hikConfig.Ip = strings.Replace(hikConfig.Ip, "https://", "", -1)
	if v, ok := h.hikMap.Load(hikConfig.Ip); !ok {
		return errors.New("没有该服务")
	} else {
		hikClient := &HikClient{}
		if hikClient, err = newHikClinet(hikConfig); err != nil {
			return err
		}
		// 结束之前的长连接
		h.DeleteHikGateway(v.(*HikClient).hikConfig.Ip)
		h.hikMap.Store(hikConfig.Ip, &hikClient)
		hikClient.ctx, hikClient.cancel = context.WithCancel(context.Background())
		go h.handelBoardcast(hikClient)
		// h.longMap.Store(hikConfig.Ip, cancel)
	}

	return
}

// 转发服务
func (h *HikGateway) transferService(key string) (client *HikClient, err error) {
	if v, ok := h.hikMap.Load(key); !ok {
		return nil, errors.New("没有该服务")
	} else {
		return v.(*HikClient), nil
	}
}

// 获取设备连接状态信息
func (h *HikGateway) GetDeviceStatus(key string) (resp []GateWayConnect) {
	h.hikMap.Range(func(key, value any) bool {
		data := GateWayConnect{}
		data.HikConfig = value.(*HikClient).hikConfig
		data.IsConnect = value.(*HikClient).isConnect
		resp = append(resp, data)
		return true
	})
	return
}

// watch全部的广播服务
func (h *HikGateway) handelBoardcast(client *HikClient) {

	if !client.isConnect {
		return
	}
	file, _ := os.OpenFile(fmt.Sprintf("./%s.txt", client.hikConfig.Ip), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()

	client.newHikClientLongConnect(func(buf Result) bool {
		file.WriteString("接收来自网关的数据\n")
		select {
		case <-client.ctx.Done():
			{
				file.WriteString("接收来自网关的数据done\n")
				client.longConnect = false
				return false
			}
		default:
			{
				file.WriteString("进入处理网关的数据default\n")
				data := Msg{}
				data.Ip = client.hikConfig.Ip
				data.Type = SUCCESS
				data.Data = buf
				client.longConnect = true
				h.broadClient.Range(func(key, value any) bool {
					file.WriteString("进入处理网关的数据default   发送到通道开始\n")
					key.(chan Msg) <- data
					file.WriteString(fmt.Sprintf("发送数据向网关注册的服务 发送到通道结束\n=======%s\n%s\n=====\n", buf.Data, buf.Type))
					return true
				})
				// os.Exit(0)
				file.WriteString("\n开始发送结束")
			}
		}
		return true
	})

}

// cancel广播服务
func (h *HikGateway) cancelBoardcast(key string) {
	if v, ok := h.hikMap.Load(key); ok {
		v.(*HikClient).cancel()
	}
}

// 注册广播客户端
func (h *HikGateway) RegisterBroadClient(channel chan Msg) {
	h.broadClient.Store(channel, nil)
}

// 删除广播客户端
func (h *HikGateway) DeleteBroadClient(channel chan Msg) {
	h.broadClient.Delete(channel)
}

// 获取网关下的所有设备
func (h *HikGateway) GetGatewayDeviceList() (resp []HikInfo) {
	h.hikMap.Range(func(key, value any) bool {
		resp = append(resp, HikInfo{
			Ip:          value.(*HikClient).hikConfig.Ip,
			Port:        value.(*HikClient).hikConfig.Port,
			Username:    value.(*HikClient).hikConfig.Username,
			Password:    value.(*HikClient).hikConfig.Password,
			IsConnect:   value.(*HikClient).isConnect,
			LongConnect: value.(*HikClient).longConnect,
		})
		return true
	})
	return
}
