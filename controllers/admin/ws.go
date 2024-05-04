// ws服务端
package admin

import (
	"encoding/json"
	"kaiyun/consts"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gorilla/websocket"
)

type WSController struct {
	AdminBaseController
}

// 接收消息结构体
type receiveMessage struct {
	Action string
	//Token  string
}

// 返回消息结构体
type returnMessage struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Action string `json:"action"`
	//Token  string      `json:"token"`
	Data interface{} `json:"data"`
}

var (
	conn     *websocket.Conn
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (c *WSController) Index() {
	var err error

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
		_, data, err := conn.ReadMessage()
		if err != nil {
			//logs.Error(err)
			//return //退出循环，并且代码不会再执行后面的语句
			goto ERR
		}
		//处理接收到消息
		if string(data) == "~H#C~" { //接收到客户端心跳包
			//回复一个心跳包
			conn.WriteMessage(websocket.TextMessage, []byte("~H#S~"))
		} else {
			var revMsg receiveMessage
			json.Unmarshal(data, &revMsg)

			//参数验证
			valid := validation.Validation{}
			valid.Required(revMsg.Action, "action")
			//valid.Required(revMsg.Token, "token")
			if valid.HasErrors() {
				conn.WriteMessage(websocket.TextMessage, []byte("invalid request, received->"+string(data)))
			}

			if revMsg.Action == "say_hi" {
				//组装数据
				resp := make(map[string]interface{}) //创建1个空集合
				c.Success(revMsg.Action, "success", resp)
			}
		}
	}

ERR:
	logs.Error(err)
	conn.Close()
}

// 成功返回
func (c *WSController) Success(action string, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	if data == nil || data == "" {
		data = struct{}{}
	}
	res := &returnMessage{
		consts.SUCCESS, msg, action, data, //0=成功
	}
	jsonString, _ := json.Marshal(res)
	conn.WriteMessage(websocket.TextMessage, []byte(jsonString)) //发送消息
}

// 失败返回
func (c *WSController) Error(action string, code int, msg string, data interface{}) {
	if code >= 0 {
		code = consts.UNKNOWN_ERROR_STATUS
	}
	if msg == "" {
		msg = "error"
	}
	if data == nil || data == "" {
		data = struct{}{}
	}
	res := &returnMessage{
		code, msg, action, data, //0=成功
	}
	jsonString, _ := json.Marshal(res)
	conn.WriteMessage(websocket.TextMessage, []byte(jsonString)) //发送消息
}
