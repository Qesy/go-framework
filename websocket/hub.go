package websocket

import (
	"github.com/Qesy/qesygo"
	"golang.org/x/net/websocket"
)

// Client 客户端结构
type Client struct {
	conn        *websocket.Conn
	uid         string
	connectTime int64
	lastTime    int64
	remoteIP    string
	ping        string
}

// ClientMsg 消息结构体
type ClientMsg struct {
	uid string
	msg string
}

// Hub 结构
type Hub struct {
	clients    map[string]*Client
	broadcast  chan string
	register   chan *Client
	unregister chan *Client
	clientMsg  chan *ClientMsg
}

// HubRouter 路由
var HubRouter = newHub()

func init() {
	go HubRouter.run()
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clientMsg:  make(chan *ClientMsg),
	}
}

func (h *Hub) run() {
	//timer1 := time.NewTicker(1 * time.Second)

	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.uid]; ok {
				if client.conn != h.clients[client.uid].conn {
					retArr := make(map[string]interface{})
					retArr["Act"] = "User_RemoteOnline"
					retArr["Code"] = 10409
					msgByte, _ := qesygo.JsonEncode(retArr)
					HubSend(h.clients[client.uid].conn, string(msgByte))
					h.clients[client.uid].conn.Close()
				}
			}
			h.clients[client.uid] = client

		case client := <-h.unregister:
			if _, ok := h.clients[client.uid]; ok {
				delete(h.clients, client.uid)
			}
			client.conn.Close()
			connectTime := qesygo.Time("Microsecond") - client.connectTime
			if connectTime > 1000 {
				connectTime = connectTime / 1000

			}
		case message := <-h.broadcast:
			for _, usClient := range h.clients {
				HubSend(usClient.conn, message)
			}
		case clientMsg := <-h.clientMsg:
			if Client, ok := h.clients[clientMsg.uid]; ok {
				HubSend(Client.conn, clientMsg.msg)
			}

			// case <-timer1.C:
			// 	qesygo.Println(qesygo.Time("Second"))
		}

	}
}
