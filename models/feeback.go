package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Feeback struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Name       string `orm:"size(50);default();null;index;description(姓名)" json:"name"`
	Mobile     string `orm:"size(20);default();null;index;description(手机号)" json:"mobile"`
	Email      string `orm:"size(100);default();null;index;description(信箱)" json:"email"`
	Feeback    string `orm:"type(text);null;description(內容)" json:"content"`
	CreateTime int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(Feeback))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Feeback) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "feeback"
}

// 获取全部列表
func (m *Feeback) All(query dto.FeebackQuery) (list []*Feeback) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Feeback))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Feeback) PageList(query dto.FeebackQuery) ([]*Feeback, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Feeback))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*Feeback
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Feeback, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Feeback) GetById(id int) (v *Feeback, err error) {
	o := orm.NewOrm()
	v = &Feeback{}
	err = o.QueryTable(new(Feeback)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Feeback) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Feeback) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Feeback) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Feeback) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Feeback) BatchAdd(data []*Feeback) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
