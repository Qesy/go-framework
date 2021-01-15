package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	wsController "todo/websocket"

	"github.com/Qesy/qesygo"
	"golang.org/x/net/websocket"
	//"github.com/golang/net/websocket"
)

// Entry 结构
type Entry struct {
	W          http.ResponseWriter
	R          *http.Request
	Controller string
	Action     string
	URL        string
	Params     []string
}

// Fetch 路由匹配
func (e *Entry) fetch() {
	if e.URL == "" {
		return
	}
	urlArr := strings.Split(e.URL, "/")
	if len(urlArr) <= 1 {
		e.Controller = urlArr[0]
	} else {
		e.Controller = urlArr[0]
		if urlArr[1] != "" {
			e.Action = urlArr[1]
		}
		e.Params = urlArr[2:len(urlArr)]
	}
}

// Run 启动所有服务
func (e *Entry) Run() {
	e.fetch()
	if e.Controller == "static" { // 静态文件服务器
		qesygo.Println(e.URL)
		http.ServeFile(e.W, e.R, e.URL)
		return
	}
	if e.Controller == "ws" { // websocket
		e.R.Header.Set("Origin", "file://")
		handle := websocket.Handler(wsController.Echo)
		handle.ServeHTTP(e.W, e.R)
		return
	}
	t := reflect.TypeOf(e)
	action := strings.Title(e.Controller) + "_" + e.Action
	fmt.Println("HTTPRECEIVE:" + action)
	m, ok := t.MethodByName(action)
	if !ok {
		fmt.Println(ok, e.URL)
		e.ErrorStatus(404)
		return
	}
	m.Func.Call([]reflect.Value{reflect.ValueOf(e)})
}

// ErrorStatus 错误状态
func (e *Entry) ErrorStatus(code int) {
	e.W.WriteHeader(code)
	fmt.Fprintf(e.W, "%d page not found", code)
}

// Error 返回数据
func (e *Entry) Error(code int) {
	retArr := make(map[string]interface{})
	retArr["act"] = e.Controller + "/" + e.Action
	retArr["code"] = code
	retArr["data"] = make(map[string]string)
	jsonByte, _ := qesygo.JsonEncode(retArr)
	fmt.Fprintf(e.W, string(jsonByte))
}

// Show 展示数据
func (e *Entry) Show(ret interface{}) {
	var str string
	switch ret.(type) {
	case []byte:
		str = fmt.Sprintf("%s", ret)
	case string:
		str = fmt.Sprintf("%s", ret)
	case map[string]interface{}:
		json, err := qesygo.JsonEncode(ret)
		if err != nil {
			return
		}
		str = fmt.Sprintf("%s", json)
	default:
		return
	}
	fmt.Fprintf(e.W, str)
}
