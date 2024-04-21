package models

import (
	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type UserLog struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Uid        string `orm:"size(32);default();null;description(用户id)" json:"uid"`
	Username   string `orm:"size(40);default();null;index;description(帐号)" json:"username"`
	Agent      string `orm:"size(200);default();null;description(客户端消息)" json:"agent"`
	Ip         string `orm:"size(15);default();null;description(客户端ip)" json:"ip"`
	Url        string `orm:"size(100);default();null;description(请求url)" json:"url"`
	Param      string `orm:"type(text);null;description(业务信息)" json:"param"`
	Title      string `orm:"size(50);null;index;description(操作说明)" json:"title"`
	CreateTime int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(UserLog))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *UserLog) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "user_log"
}
