package models

import (
	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type RolePermission struct {
	Id            int    `orm:"pk;auto;default();description(ID)" json:"id"`
	RoleId        int    `orm:"default(0);null;description(角色id)" json:"role_id"`
	PeermissionId int    `orm:"default(0);null;description(权限id)" json:"permission_id"`
	CreateTime    int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser    string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime    int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser    string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime    int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser    string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(RolePermission))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *RolePermission) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "role_permission"
}
