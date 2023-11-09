package hk_gateway

// 发送校验账号密码
func (h *HikGateway) UserCheck(key string) (result *UserCheckRes, err error) {
	req, err := h.transferService(key)
	if err != nil {
		return
	}
	return req.UserCheck()
}

// /ISAPI/AccessControl/UserInfo/Record?format=json
func (h *HikGateway) AddPersonInfo(key string, person AddPersonInfoReq) (*ErrorMsg, error) {
	req, err := h.transferService(key)
	if err != nil {
		return nil, nil
	}
	return req.AddPersonInfo(person)
}

// 设置人脸信息
// /ISAPI/Intelligent/FDLib/FDSetUp
func (c *HikGateway) SetFaceInfo(key string, reqBody SetFaceInfo) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.SetFaceInfo(reqBody)
}

// 修改人员信息
// ISAPI/AccessControl/UserInfo/Modify?format=json
func (c *HikGateway) ModifyPersonInfo(key string, person AddPersonInfoReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.ModifyPersonInfo(person)
}

// 删除人员信息
// ISAPI/AccessControl/UserInfo/Delete?format=json
func (c *HikGateway) DeletePersonInfo(key string, reqBody DelPersonInfo) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.DeletePersonInfo(reqBody)
}

// 获取人员列表
// /ISAPI/AccessControl/UserInfo/Search
func (c *HikGateway) GetPersonInfoList(key string, reqBody GetPersonInfoList) (result *UserInfoSearch, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetPersonInfoList(reqBody)
}

// 日志查询
// /ISAPI/ContentMgmt/logSearch
func (c *HikGateway) LogSearch(key string, reqBody LogSearch) (result *LogSearchRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.LogSearch(reqBody)
}

// 获取人员数量信息
// /ISAPI/AccessControl/UserInfo/Count
func (c *HikGateway) GetPersonInfoCount(key string) (result *GetPersonInfoCountRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetPersonInfoCount()
}

// 查询设备中已有的人脸数量及人脸信息
// /ISAPI/Intelligent/FDLib/Count?format=json
func (c *HikGateway) GetFaceInfoCount(key string) (result *GetFaceInfoCountRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetFaceInfoCount()
}

// 查询指定或全部人员的卡数量
// /ISAPI/AccessControl/CardInfo/Count?format=json
func (c *HikGateway) GetCardInfoCount(key string) (result *GetCardInfoCountRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetCardInfoCount()
}

// 查询门禁事件条数。
// /ISAPI/AccessControl/AcsEventTotalNum?format=json
func (c *HikGateway) GetAcsEventTotalNum(key string, reqBody GetAcsEventTotalNumReq) (result *GetAcsEventTotalNumRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetAcsEventTotalNum(reqBody)
}

// 获取门禁主机工作状态
// /ISAPI/AccessControl/AcsWorkStatus?format=json
func (c *HikGateway) GetAcsWorkStatus(key string) (result *GetAcsWorkStatusRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetAcsWorkStatus()
}

// 添加卡信息
// /ISAPI/AccessControl/CardInfo/Record
func (c *HikGateway) AddCardInfo(key string, reqBody AddCardInfoReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.AddCardInfo(reqBody)
}

// 获取用户的图片
func (c *HikGateway) GetPicture(key string, reqBody GetPictureReq) (result []byte, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}

	return req.GetPicture(reqBody)
}

// 删除卡信息
// /ISAPI/AccessControl/CardInfo/Delete?format=json
func (c *HikGateway) DeleteCardInfo(key string, reqBody DelCardInfoReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.DeleteCardInfo(reqBody)
}

// 获取卡信息
// /ISAPI/AccessControl/CardInfo/Search?format=json
func (c *HikGateway) GetCardInfo(key string, reqBody GetCardInfoReq) (result *GetCardInfoRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetCardInfo(reqBody)
}

// 查询门禁事件
// /ISAPI/AccessControl/AcsEvent?format=json
func (c *HikGateway) GetAcsEvent(key string, reqBody GetAcsEventReq) (result *GetAcsEventRes, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetAcsEvent(reqBody)
}

// 配置门禁主机参数
// /ISAPI/AccessControl/AcsCfg?format=json
func (c *HikGateway) SetAcsCfg(key string, reqBody SetAcsCfgReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.SetAcsCfg(reqBody)
}

// /ISAPI/AccessControl/AcsEvent/StorageCfg?format=json
func (c *HikGateway) SetStorageCfg(key string, reqBody SetStorageCfgReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.SetStorageCfg(reqBody)
}

// 配置人员及凭证显示参数
// /ISAPI/AccessControl/userAndRightShow?format=json
func (c *HikGateway) GetUserAndRightShow(key string, reqBody GetUserAndRightShowReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetUserAndRightShow(reqBody)
}

// 远程控门
// PUT/ISAPI/AccessControl/RemoteControl/door/<doorID>
func (c *HikGateway) RemoteControlDoor(key string, reqBody RemoteControlDoorReq) (result *ErrorMsg, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.RemoteControlDoor(reqBody)
}

// 获取设备信息参数
// GET/ISAPI/System/deviceInfo
func (c *HikGateway) GetDeviceInfo(key string) (result *DeviceInfo, err error) {
	req, err := c.transferService(key)
	if err != nil {
		return
	}
	return req.GetDeviceInfo()
}
