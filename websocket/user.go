package websocket

// UserLoginReq 用户登录
func (e *Entry) UserLoginReq() {
	retArr := make(map[string]interface{})
	retArr["Act"] = e.Act
	retArr["Data"] = e.Data
	e.Send(retArr)
	UID := e.Data["UID"]
	e.Register(UID)
}

// UserMsgReq 发消息测试
func (e *Entry) UserMsgReq() {
	retArr := make(map[string]interface{})
	retArr["Act"] = e.Act
	retArr["Data"] = e.Data
	e.Send(retArr)
}

// UserLogoutReq 用户退出
func (e *Entry) UserLogoutReq() {
	retArr := make(map[string]interface{})
	retArr["Act"] = e.Act
	retArr["Data"] = e.Data
	e.Send(retArr)
	e.Unregister()
}
