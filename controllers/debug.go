package controllers

import (

	//wsController "ball_game_server/websocket"

	"github.com/Qesy/qesygo"

	//"math"

	"time"
)

func (e *Entry) Debug_platform() {
	// aa, _ := models.PlatformGet()
	// qesygo.Println(aa)
}

func (e *Entry) Debug_test() {
	/*a := &wsController.HubRouter
	a.broadcast <- "fsfd"*/
	//wsController.HubRouter.broadcast <- "fuck"
	/*t0 := time.Now()
	qesygo.Println(t0)*/
	timer1()
	//wsController.Broadcast("aaaa")
}

func timer1() { //-- 定时器 --
	timer1 := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-timer1.C:
			qesygo.Println("执行")
		}
	}
}

func timer2() { //-- 多少时间后执行 --
	/*f := func() {
		qesygo.Println("Time out")
	}
	time.AfterFunc(5*time.Second, f)*/
	time1 := time.After(5 * time.Second)

	for {
		/*select {
		case <-timer1.C:
			qesygo.Println("执行")
		}*/
		select {
		/*case m := <-timer.c:
		handle(m)*/
		case <-time1:
			qesygo.Println("timed out")
		}
	}

}

// func (e *Entry) Debug_pwd() {
// 	var m QesyDb.Model
// 	result, _ := m.SetTable("user").SetField("UserId,Password").ExecSelect()
// 	for _, v := range result {
// 		m.SetTable("user").SetUpdate(map[string]string{"Password": lib.Md5(v["Password"])}).SetWhere(map[string]string{"UserId": v["UserId"]}).ExecUpdate()
// 	}
// 	return
// }
