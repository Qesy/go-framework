package websocket

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/Qesy/qesygo"
	"golang.org/x/net/websocket"
)

// Entry 结构
type Entry struct {
	Conn *websocket.Conn
	UID  string
	Act  string
	Data map[string]string
}

// MsgStr 消息结构体
type MsgStr struct {
	Act  string
	Data map[string]string
}

// Sys 获取系统配置
var Sys map[string]string

// Deadline Ping值
var Deadline = 1000

func init() {
	// var err error
	// Sys, err = models.SysGetConf()
	// if err != nil {
	// 	qesygo.Die("SysGetConf")
	// }
	// fmt.Println("websocket connect sueccss")
}

// Echo 运行
func Echo(conn *websocket.Conn) {
	e := &Entry{Conn: conn}
	e.SetDeadline()
	i := 0
	for {
		i++
		qesygo.Println(i)
		var reply string
		if err := websocket.Message.Receive(e.Conn, &reply); err != nil {
			fmt.Println("Receive Err : ", err)
			e.Unregister()
			break
		}
		var replyArr MsgStr
		if err := qesygo.JsonDecode([]byte(reply), &replyArr); err != nil {
			fmt.Println("Receive Err : ", err)
			e.SendError(401, err.Error())
			continue
		}
		RecPrint("\033[35;4mSOCKETREC:(", "Act:"+replyArr.Act+" Data:", replyArr.Data, ")\033[0m")
		e.Act, e.Data = replyArr.Act, replyArr.Data

		if e.UID == "" && e.Act != "UserLoginReq" {
			e.SendError(402, "No Login")
			fmt.Println("No Login ")
			break
		}
		t := reflect.TypeOf(e)
		if m, ok := t.MethodByName(strings.Title(e.Act)); ok {
			m.Func.Call([]reflect.Value{reflect.ValueOf(e)})
		} else {
			e.SendError(403, "MethodByName Failed")
		}
		continue
	}
}

// SetDeadline 设置PING值
func (e *Entry) SetDeadline() *Entry {
	t := time.Now().Add(time.Duration(Deadline) * time.Second)
	e.Conn.SetDeadline(t)
	return e
}

// Register 用户唯一登录口
func (e *Entry) Register(UserID string) {
	connectTime := qesygo.Time("Microsecond")
	remoteIP, _, _ := net.SplitHostPort(e.Conn.Request().RemoteAddr)
	client := Client{uid: UserID, conn: e.Conn, connectTime: connectTime, remoteIP: remoteIP}
	e.UID = UserID
	HubRouter.register <- &client
	fmt.Println("Login:", client)
}

// Unregister 用户退出
func (e *Entry) Unregister() {
	client := HubRouter.clients[e.UID]
	fmt.Println("LogOut:", e.UID)
	if client != nil && e.Conn == client.conn {
		HubRouter.unregister <- client
	}
}

// Ping 过期时间
func (e *Entry) Ping() {
	e.SetDeadline()
	data := make(map[string]interface{})
	data["t"] = qesygo.TimeStr("Second")
	retArr := make(map[string]interface{})
	retArr["act"] = e.Act
	retArr["data"] = data
	retArr["code"] = 200
	e.Send(retArr)

}

// SendError 发送错误
func (e *Entry) SendError(ErrCode int32, msg string) *Entry {
	retArr := make(map[string]interface{})
	retArr["act"] = e.Act
	retArr["data"] = e.Data
	retArr["msg"] = msg
	retArr["code"] = ErrCode
	e.Send(retArr)
	return e
}

// SendSueccss 发送成功数据
func (e *Entry) SendSueccss(msg string) *Entry {
	retArr := make(map[string]interface{})
	retArr["act"] = e.Act
	retArr["data"] = e.Data
	retArr["code"] = 200
	e.Send(retArr)
	return e
}

// Send 发送消息
func (e *Entry) Send(msgArr map[string]interface{}) *Entry {
	msgByte, _ := qesygo.JsonEncode(msgArr)
	Print("Single", msgByte, []string{e.UID})
	HubSend(e.Conn, string(msgByte))
	return e
}

// Send 发送消息
func Send(UserID string, msgArr map[string]interface{}) {
	msgByte, _ := qesygo.JsonEncode(msgArr)
	Print("Single", msgByte, []string{UserID})
	HubRouter.clientMsg <- &ClientMsg{uid: UserID, msg: string(msgByte)}
}

// SendMultiple 指定多发
func SendMultiple(IDArr []string, msgArr map[string]interface{}) {
	msgByte, _ := qesygo.JsonEncode(msgArr)
	Print("Multiple", msgByte, IDArr)
	for _, v := range IDArr {
		HubRouter.clientMsg <- &ClientMsg{uid: v, msg: string(msgByte)}
	}
}

// HubSend 底层发送
func HubSend(conn *websocket.Conn, msg string) error {
	return websocket.Message.Send(conn, msg)
}

// Broadcast 广播
func Broadcast(msgArr map[string]interface{}) {
	msgByte, _ := qesygo.JsonEncode(msgArr)
	Print("Broadcast", msgByte, []string{})
	HubRouter.broadcast <- string(msgByte)
}

// Print 发送打印
func Print(SendType string, Msg []byte, userIDArr []string) {
	UIDStr := qesygo.Implode(userIDArr, ",")
	switch SendType {
	case "Single", "Multiple":
		fmt.Println("\033[36;4mSOCKETSEND:(", "Type:"+SendType, ", UidArr:"+UIDStr, ")"+" \nData:", string(Msg)+"\033[0m")
	case "BroadcastAll":
		fmt.Println("\033[36;4mSOCKETSEND:(", "Type:"+SendType, ")"+" \nData:", string(Msg)+"\033[0m")
	}
}

// RecPrint 接收打印
func RecPrint(str string, str2 string, str3 map[string]string, str4 string) {
	//return
	fmt.Println(str, str2, str3, str4)
}

// UserCount 统计用户
func UserCount() int {
	return len(HubRouter.clients)
}
