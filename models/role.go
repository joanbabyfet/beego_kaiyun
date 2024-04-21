package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Role struct {
	Id          int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Pid         int    `orm:"default(0);null;description(上级id)" json:"pid"`
	Code        string `orm:"size(20);default();null;description(编码)" json:"code"`
	Name        string `orm:"size(50);default();null;index;description(角色名称)" json:"name"`
	Description string `orm:"default();null;description(描述)" json:"description"`
	Sort        int16  `orm:"default(0);null;description(排序: 数字小的排前面)" json:"sort"`
	Status      int8   `orm:"default(1);null;description(状态: 0=禁用 1=启用)" json:"status"`
	CreateTime  int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser  string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime  int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser  string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime  int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser  string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(Role))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Role) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "role"
}

// 获取全部列表
func (m *Role) All(query dto.RoleQuery) (list []*Role) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Role) PageList(query dto.RoleQuery) ([]*Role, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*Role
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Role, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Role) GetById(id int) (v *Role, err error) {
	o := orm.NewOrm()
	v = &Role{}
	err = o.QueryTable(new(Role)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Role) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Role) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Role) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Role) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Role) BatchAdd(data []*Role) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
