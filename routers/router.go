// 路由文件
package routers

import (
	admin "kaiyun/controllers/admin"
	controllers "kaiyun/controllers/web"

	beego "github.com/beego/beego/v2/server/web"
)

// 初始化
func init() {
	//官网前台
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/case", &controllers.MainController{}, "get:Case")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	beego.Router("/partner", &controllers.MainController{}, "get:Partner")
	beego.Router("/product", &controllers.MainController{}, "get:Product")
	beego.Router("/quality", &controllers.MainController{}, "get:Quality")
	beego.Router("/service", &controllers.MainController{}, "get:Service")
	beego.Router("/content/:id", &controllers.MainController{}, "get:Content")
	beego.Router("/search", &controllers.MainController{}, "get:Search")
	beego.Router("/ws_test", &controllers.MainController{}, "get:WSTest")

	//ws服务端
	beego.Router("/ws", &admin.WSController{}, "get:Index")

	//后端接口
	admin := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/article", &admin.ArticleController{}, "get:Index"),
			beego.NSRouter("/article/detail", &admin.ArticleController{}, "get:Detail"),
			beego.NSRouter("/article/save", &admin.ArticleController{}, "post:Save"),
			beego.NSRouter("/article/delete", &admin.ArticleController{}, "post:Delete"),
			beego.NSRouter("/article/enable", &admin.ArticleController{}, "post:Enable"),
			beego.NSRouter("/article/disable", &admin.ArticleController{}, "post:Disable"),
			beego.NSRouter("/ad", &admin.AdController{}, "get:Index"),
			beego.NSRouter("/ad/detail", &admin.AdController{}, "get:Detail"),
			beego.NSRouter("/ad/save", &admin.AdController{}, "post:Save"),
			beego.NSRouter("/ad/delete", &admin.AdController{}, "post:Delete"),
			beego.NSRouter("/ad/enable", &admin.AdController{}, "post:Enable"),
			beego.NSRouter("/ad/disable", &admin.AdController{}, "post:Disable"),
			beego.NSRouter("/upload", &admin.UploadController{}, "post:Upload"),
			beego.NSRouter("/download", &admin.UploadController{}, "get:Download"),
			beego.NSRouter("/captcha", &admin.CommonController{}, "get:Captcha"), //获取验证码
			beego.NSRouter("/reload_captcha", &admin.CommonController{}, "get:Captcha"),
			beego.NSRouter("/chat_gpt", &admin.CommonController{}, "get:ChatGPT"),
			beego.NSRouter("/ip", &admin.CommonController{}, "get:Ip"),
			beego.NSRouter("/ping", &admin.CommonController{}, "get:Ping"),
			beego.NSRouter("/test", &admin.TestController{}, "get:Test"),
		),
	)
	beego.AddNamespace(admin)
}
