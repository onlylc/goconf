package ws

// import (
// 	"context"
// 	"fmt"
// 	"goconf/core/sdk/pkg"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// )

// type Manager struct {
// 	Group                   map[string]map[string]*Client
// 	groupCount, clientCount uint
// 	Lock                    sync.Mutex
// 	Register, UnRegister    chan *Client
// 	Message                 chan *MessageData
// 	GroupMessage            chan *GroupMessageData
// 	BroadCastMessage        chan *BroadCastMessageData
// }

// type Client struct {
// 	Id, Group  string
// 	Context    context.Context
// 	CancelFunc context.CancelFunc
// 	Socket     *websocket.Conn
// 	Message    chan []byte
// }

// type MessageData struct {
// 	Id, Group string
// 	Context   context.Context
// 	Message   []byte
// }

// type GroupMessageData struct {
// 	Group   string
// 	Message []byte
// }

// type BroadCastMessageData struct {
// 	Message []byte
// }

// func (manager *Manager) Start() {
// 	log.Printf("websocket manage start")
// 	for {
// 		select {
// 		case client := <-manager.Register:
// 			log.Printf("client [%s] connect", client.Id)
// 			log.Printf("register client [%s] to group [%s]", client.Id, client.Group)

// 			manager.Lock.Lock()
// 			if manager.Group[client.Group] == nil {
// 				manager.Group[client.Group] = make(map[string]*Client)
// 				manager.groupCount += 1
// 			}
// 			manager.Group[client.Group][client.Id] = client
// 			manager.clientCount += 1
// 			manager.Lock.Unlock()
// 		case client := <-manager.UnRegister:
// 			log.Printf("unregister client [%s] from group [%s]", client.Id, client.Group)
// 			manager.Lock.Lock()
// 			if mGroup, ok := manager.Group[client.Group]; ok {
// 				if mClient, ok := mGroup[client.Id]; ok {
// 					close(mClient.Message)
// 					delete(mGroup, client.Id)
// 					manager.clientCount -= 1
// 					if len(mGroup) == 0 {
// 						delete(manager.Group, client.Group)
// 						manager.groupCount -= 1
// 					}
// 					mClient.CancelFunc()
// 				}
// 			}
// 			manager.Lock.Unlock()
// 		}

// 	}

// }

// // 处理单个 client 发送数据
// func (manager *Manager) SendService() {
// 	for {
// 		select {
// 		case data := <-manager.Message:
// 			if groupMap, ok := manager.Group[data.Group]; ok {
// 				if conn, ok := groupMap[data.Id]; ok {
// 					conn.Message <- data.Message
// 				}
// 			}
// 		}
// 	}
// }

// func (manager *Manager) SendAllService() {
// 	for {
// 		select {
// 		case data := <-manager.BroadCastMessage:
// 			for _, v := range manager.Group {
// 				for _, conn := range v {
// 					conn.Message <- data.Message
// 				}
// 			}
// 		}
// 	}
// }

// var WebsocketManager = Manager{
// 	Group:            make(map[string]map[string]*Client),
// 	Register:         make(chan *Client, 128),
// 	UnRegister:       make(chan *Client, 128),
// 	GroupMessage:     make(chan *GroupMessageData, 128),
// 	Message:          make(chan *MessageData, 128),
// 	BroadCastMessage: make(chan *BroadCastMessageData, 128),
// 	groupCount:       0,
// 	clientCount:      0,
// }

// func (manager *Manager) WsClient(c *gin.Context) {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	upGrader := websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
// 	}

// 	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Printf("websocket connect error: %s", c.Param("channel"))
// 		return
// 	}

// 	fmt.Println("token: ", c.Query("token"))

// 	client := &Client{
// 		Id:         c.Param("id"),
// 		Group:      c.Param("channel"),
// 		Context:    ctx,
// 		CancelFunc: cancel,
// 		Socket:     conn,
// 		Message:    make(chan []byte, 1024),
// 	}
// 	// manager.RegisterClient(client)
// 	// go c.Client.Read(ctx)
// 	// go client.Write(ctx)
// 	time.Sleep(time.Second * 15)

// 	pkg.FileMonitoringById(ctx, "temp/logs/job/db-20200820", c.Param("id"), c.Param("channel"),SendOne)
// }

// func (manager *Manager) UnWsClient(c *gin.Context) {
// 	id := c.Param("id")
// 	group := c.Param("channel")
// 	// WsLoggout(id, group)
// 	c.Set("result", "ws close succes")
// 	c.JSON(http.StatusOK,gin.H{
// 		"code": http.StatusOK,
// 		"data": "ws close succes",
// 		"msg": "success",
// 	})
// }