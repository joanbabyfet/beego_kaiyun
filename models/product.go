package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Product struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Catid      int    `orm:"default(0);null;description(分类id 1产品 2合作伙伴 3行业客户)" json:"catid"`
	Title      string `orm:"size(50);default();null;index;description(標題)" json:"title"`
	Info       string `orm:"default();null;description(简介)" json:"info"`
	Content    string `orm:"type(text);null;description(內容)" json:"content"`
	Img        string `orm:"size(100);default();null;description(圖片)" json:"img"`
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
	orm.RegisterModel(new(Product))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Product) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "product"
}

// 获取全部列表
func (m *Product) All(query dto.ProductQuery) (list []*Product) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
	qs = qs.Filter("delete_time", 0) //未删除
	if query.Limit != 0 {
		qs = qs.Limit(query.Limit)
	}
	if !reflect.ValueOf(&query.Status).IsNil() {
		qs = qs.Filter("status", query.Status)
	}
	if query.Catid != 0 {
		qs = qs.Filter("catid", query.Catid)
	}
	if len(query.Catids) > 0 {
		qs = qs.Filter("catid__in", query.Catids)
	}
	if len(query.Title) > 1 {
		qs = qs.Filter("title__icontains", query.Title)
	}
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Product) PageList(query dto.ProductQuery) ([]*Product, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
	qs = qs.Filter("delete_time", 0) //未删除
	if query.Limit != 0 {
		qs = qs.Limit(query.Limit)
	}
	if !reflect.ValueOf(&query.Status).IsNil() {
		qs = qs.Filter("status", query.Status)
	}
	if query.Catid != 0 {
		qs = qs.Filter("catid", query.Catid)
	}
	if len(query.Catids) > 0 {
		qs = qs.Filter("catid__in", query.Catids)
	}
	if len(query.Title) > 1 {
		qs = qs.Filter("title__icontains", query.Title)
	}
	//总条数
	count, _ := qs.Count()
	var list []*Product
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Product, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Product) GetById(id int) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{}
	err = o.QueryTable(new(Product)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Product) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Product) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Product) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Product) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Product) BatchAdd(data []*Product) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
