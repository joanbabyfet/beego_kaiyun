package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Permission struct {
	Id          int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Pid         int    `orm:"default(0);null;description(上级id)" json:"pid"`
	Name        string `orm:"size(50);default();null;index;description(权限名)" json:"name"`
	Description string `orm:"default();null;description(描述)" json:"description"`
	Url         string `orm:"size(100);default();null;description(访问路径)" json:"url"`
	Perms       string `orm:"size(100);default();null;description(权限码)" json:"perms"`
	Type        int16  `orm:"default(0);null;description(类型)" json:"type"`
	Icon        string `orm:"size(100);default();null;description(图标)" json:"icon"`
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
	orm.RegisterModel(new(Permission))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Permission) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "permission"
}

// 获取全部列表
func (m *Permission) All(query dto.PermissionQuery) (list []*Permission) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Permission) PageList(query dto.PermissionQuery) ([]*Permission, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*Permission
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Permission, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Permission) GetById(id int) (v *Permission, err error) {
	o := orm.NewOrm()
	v = &Permission{}
	err = o.QueryTable(new(Permission)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Permission) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Permission) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Permission) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Permission) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Permission) BatchAdd(data []*Permission) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
