// 父控制器
package controllers

import (
	"kaiyun/consts"
	"kaiyun/utils"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
}

// 全局变量
var (
	//cpt         *captcha.Captcha
	lang        = "lang"
	types       = "types"
	accept_lang = "language"
)

// 初始化, 先于Prepare函数
func init() {
	setLocale() // 设置多语言文件

	//初始化验证码, 放在Prepare函数会报错
	// store := cache.NewMemoryCache()
	// cpt = captcha.NewWithFilter("/captcha/", store)
	// cpt.ChallengeNums = 4
	// cpt.StdWidth = 100
	// cpt.StdHeight = 40

	//初始化表单验证信息
	utils.SetVerifyMessage()
}

// 定义prepare方法, 用户扩展用
func (c *BaseController) Prepare() {
	setLang(c) //切换语言
}

// 设置多语言文件
func setLocale() {
	SMH := "::"
	SX := "|"
	lang_types, _ := beego.AppConfig.String(lang + SMH + types)
	arr_lang_types := strings.Split(lang_types, SX)
	for _, lang := range arr_lang_types {
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("本地化文件设置失败")
			return
		}
	}
}

// 切换语言, 打印 i18n.Tr(c.Lang, "api_param_error")
func setLang(c *BaseController) {
	lang := c.GetString(lang)
	if lang == "" { //获取请求头参数 Accept-Language
		al := c.Ctx.Request.Header.Get(accept_lang)
		lang = al
	}

	if !i18n.IsExist(lang) {
		lang = "cn" //默认为中文
	}
	c.Lang = lang
}

// 封装接口统一返回json格式
type ReturnMsg struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int         `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// @Title API成功响应
// @Description API成功响应
// @Param msg 成功消息
// @Param data 成功返回信息
func (c *BaseController) SuccessJson(msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	if data == nil || data == "" {
		data = struct{}{}
	}
	timestamp := utils.Timestamp()
	res := &ReturnMsg{
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
func (c *BaseController) ErrorJson(code int, msg string, data interface{}) {
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
	res := &ReturnMsg{
		code, msg, timestamp, data,
	}
	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}
