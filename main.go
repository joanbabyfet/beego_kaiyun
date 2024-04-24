// 主入口
package main

import (
	_ "kaiyun/routers"

	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化
func init() {
	max_idle := 30 //最大空闲连接
	max_conn := 30 //最大数据库连接
	db_host, _ := beego.AppConfig.String("db_host")
	db_port, _ := beego.AppConfig.String("db_port")
	db_user, _ := beego.AppConfig.String("db_user")
	db_password, _ := beego.AppConfig.String("db_password")
	db_name, _ := beego.AppConfig.String("db_name")
	dsn := db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4"
	//注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	err := orm.RegisterDataBase("default", "mysql", dsn, max_idle, max_conn)
	if err != nil {
		logs.Error("连接数据库出错", err)
	}

	// 自动建表 (前提为数据库与model已注册)
	name := "default" //数据库别名
	force := false    //drop table 后再建表
	verbose := true   //打印执行过程
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		logs.Error(err)
	}

	//日志设置
	logs.Async() //支持非同步
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/app.log","separate":["error", "info"]}`)

	//后端解决跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true, //允许访问所有源
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"*"},
		// 允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//允许共享身份验证凭据, 例 cookie
		//AllowCredentials: true,
	}))

	// 全局开启打印sql语句
	orm.Debug = false

	//设置静态资源路径
	beego.SetStaticPath("/uploads", "uploads")

	//初始化定时任务
	//utils.CrontabInit()
	//utils.TbCrontabInit()
}

func main() {
	//执行 orm 命令 自动建表
	orm.RunCommand()

	beego.Run()
}
