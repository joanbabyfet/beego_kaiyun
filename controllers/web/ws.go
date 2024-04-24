// ws服务端
package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gorilla/websocket"
)

type WSController struct {
	BaseController
}

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (c *WSController) Index() {
	var conn *websocket.Conn
	var err error
	var data []byte

	conn, err = upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		//logs.Error(err)
		//http.NotFound(c.Ctx.ResponseWriter, c.Ctx.Request)
		//return
		goto ERR
	}

	//启动协程
	// go func() {
	// 	//主动向客户端发心跳
	// 	for {
	// 		err = conn.WriteMessage(websocket.TextMessage, []byte("~H#S~"))
	// 		if err != nil {
	// 			return //退出循环，并且代码不会再执行后面的语句
	// 		}
	// 		//心跳每1秒发送1次
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	for {
		_, data, err = conn.ReadMessage()
		if err != nil {
			//logs.Error(err)
			//return //退出循环，并且代码不会再执行后面的语句
			goto ERR
		}
		//处理接收到消息
		if string(data) == "~H#C~" { //检测心跳
			conn.WriteMessage(websocket.TextMessage, []byte("~H#S~"))
		}
	}

ERR:
	logs.Error(err)
	conn.Close()
}
