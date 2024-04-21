package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Category struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Pid        int    `orm:"default(0);null;description(上级id)" json:"pid"`
	Name       string `orm:"size(50);default();null;index;description(名称)" json:"name"`
	Type       int16  `orm:"default(0);null;description(类型1=友情链接 2资源)" json:"type"`
	Banner     string `orm:"size(100);default();null;description(图片)" json:"img"`
	Url        string `orm:"size(100);default();null;description(链接)" json:"url"`
	Sort       int16  `orm:"default(0);null;description(排序: 数字小的排前面)" json:"sort"`
	Status     int8   `orm:"default(1);null;description(状态: 0=禁用 1=启用)" json:"status"`
	CreateTime int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(Category))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Category) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "category"
}

// 获取全部列表
func (m *Category) All(query dto.CategoryQuery) (list []*Category) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Category))
	qs = qs.Filter("delete_time", 0) //未删除
	qs = qs.Filter("type", query.Type)
	if !reflect.ValueOf(&query.Status).IsNil() {
		qs = qs.Filter("status", query.Status)
	}
	if !reflect.ValueOf(&query.Pid).IsNil() {
		qs = qs.Filter("pid", query.Pid)
	}
	if !reflect.ValueOf(&query.Type).IsNil() {
		qs = qs.Filter("type", query.Type)
	}
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Category) PageList(query dto.CategoryQuery) ([]*Category, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Category))
	qs = qs.Filter("delete_time", 0) //未删除
	if !reflect.ValueOf(&query.Status).IsNil() {
		qs = qs.Filter("status", query.Status)
	}
	if !reflect.ValueOf(&query.Pid).IsNil() {
		qs = qs.Filter("pid", query.Pid)
	}
	if !reflect.ValueOf(&query.Type).IsNil() {
		qs = qs.Filter("type", query.Type)
	}
	//总条数
	count, _ := qs.Count()
	var list []*Category
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Category, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Category) GetById(id int) (v *Category, err error) {
	o := orm.NewOrm()
	v = &Category{}
	err = o.QueryTable(new(Category)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Category) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Category) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Category) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Category) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Category) BatchAdd(data []*Category) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
