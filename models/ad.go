package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Ad struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Catid      int    `orm:"default(0);null;description(分類id)" json:"catid"`
	Title      string `orm:"size(50);default();null;index;description(标题)" json:"title"`
	Img        string `orm:"size(100);default();null;description(图片)" json:"img"`
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
	orm.RegisterModel(new(Ad))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Ad) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "ad"
}

// 获取全部列表
func (m *Ad) All(query dto.AdQuery) (list []*Ad) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ad))
	qs = qs.Filter("delete_time", 0) //未删除
	if !reflect.ValueOf(&query.Status).IsNil() {
		qs = qs.Filter("status", query.Status)
	}
	if query.Type != 0 {
		cat := new(Category)
		var ids []int
		var cats []*Category
		o.QueryTable(cat).Filter("type", query.Type).All(&cats, "id")
		for _, v := range cats {
			ids = append(ids, int(v.Id))
		}
		qs = qs.Filter("catid__in", ids)
	}
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Ad) PageList(query dto.AdQuery) ([]*Ad, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ad))
	qs = qs.Filter("delete_time", 0) //未删除
	if !reflect.ValueOf(&query.Status).IsNil() {
		qs = qs.Filter("status", query.Status)
	}
	if query.Catid != 0 {
		qs = qs.Filter("catid", query.Catid)
	}
	if query.Type != 0 {
		cat := new(Category)
		var ids []int
		var cats []*Category
		o.QueryTable(cat).Filter("type", query.Type).All(&cats, "id")
		for _, v := range cats {
			ids = append(ids, int(v.Id))
		}
		qs.Filter("catid__in", ids)
	}
	//总条数
	count, _ := qs.Count()
	var list []*Ad
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Ad, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Ad) GetById(id int) (v *Ad, err error) {
	o := orm.NewOrm()
	v = &Ad{}
	err = o.QueryTable(new(Ad)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Ad) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Ad) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Ad) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Ad) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Ad) BatchAdd(data []*Ad) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
