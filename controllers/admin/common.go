package admin

import (
	"kaiyun/utils"

	"github.com/beego/beego/v2/core/validation"
)

type CommonController struct {
	AdminBaseController
}

// 获取验证码
// func (c *CommonController) Captcha() {
// 	c.TplName = "admin/captcha.html"
// }

// 获取验证码
func (c *CommonController) Captcha() {
	id, b64s, _, err := utils.GetCaptcha()
	if err != nil {
		c.ErrorJson(-1, "生成验证码错误", nil)
	}

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["key"] = id
	resp["img"] = b64s
	c.SuccessJson("success", resp)
}

// 获取列表
func (c *CommonController) ChatGPT() {
	keyword := c.GetString("keyword")

	type ChatGPT struct {
		Keyword string
	}
	chat_gpt := ChatGPT{Keyword: keyword}

	//参数验证
	valid := validation.Validation{}
	valid.Required(chat_gpt.Keyword, "keyword")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	stat, content := utils.ChatGPT(keyword)
	if !stat {
		c.ErrorJson(-1, "发送错误", nil)
	}
	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["content"] = content

	c.SuccessJson("success", resp)
}

// 返回客户端ip
func (c *CommonController) Ip() {
	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	//resp["ip"] = c.Ctx.Input.IP() //获取到的IP是内网IP不满足需求
	resp["ip"] = c.Ctx.Request.Header.Get("X-Real-ip") //nginx 中proxy_set_header 设置的值
	c.SuccessJson("success", resp)
}

// 检测用,可查看是否返回信息及时间戳
func (c *CommonController) Ping() {
	c.SuccessJson("pong", nil)
}
