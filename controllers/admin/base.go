// 父控制器
package admin

import (
	"kaiyun/consts"
	"kaiyun/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type AdminBaseController struct {
	beego.Controller
}

// 全局变量
//var cpt *captcha.Captcha

// 封装接口统一返回json格式
type ReturnMsg struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int         `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// 初始化, 先于Prepare函数
func init() {
	//初始化验证码, 放在Prepare函数会报错
	// store := cache.NewMemoryCache()
	// cpt = captcha.NewWithFilter("/captcha/", store)
	// cpt.ChallengeNums = 4
	// cpt.StdWidth = 100
	// cpt.StdHeight = 40
}

// 定义prepare方法, 用户扩展用
func (c *AdminBaseController) Prepare() {
	//c.o = orm.NewOrm()

}

// @Title API成功响应
// @Description API成功响应
// @Param msg 成功消息
// @Param data 成功返回信息
func (c *AdminBaseController) SuccessJson(msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	if data == nil || data == "" {
		data = struct{}{}
	}
	timestamp := utils.Timestamp()
	res := ReturnMsg{
		consts.SUCCESS, msg, timestamp, data, //0=成功
	}
	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}

// @Title API失败响应
// @Description API失败响应
// @Param code 错误码
// @Param msg 异常消息
// @Param data 异常返回信息
func (c *AdminBaseController) ErrorJson(code int, msg string, data interface{}) {
	if code >= 0 {
		code = consts.UNKNOWN_ERROR_STATUS
	}
	if msg == "" {
		msg = "error"
	}
	if data == nil || data == "" {
		data = struct{}{}
	}
	timestamp := utils.Timestamp()
	res := ReturnMsg{
		code, msg, timestamp, data,
	}
	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}
