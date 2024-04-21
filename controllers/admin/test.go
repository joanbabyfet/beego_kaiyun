package admin

type TestController struct {
	AdminBaseController
}

// 测试用
func (c *TestController) Test() {
	c.SuccessJson("success", nil)
}
