package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Feedback struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Name       string `orm:"size(50);default();null;index;description(姓名)" json:"name"`
	Mobile     string `orm:"size(20);default();null;index;description(手机号)" json:"mobile"`
	Email      string `orm:"size(100);default();null;index;description(信箱)" json:"email"`
	Content    string `orm:"type(text);null;description(內容)" json:"content"`
	CreateTime int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(Feedback))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Feedback) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "feedback"
}

// 获取全部列表
func (m *Feedback) All(query dto.FeedbackQuery) (list []*Feedback) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Feedback))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Feedback) PageList(query dto.FeedbackQuery) ([]*Feedback, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Feedback))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*Feedback
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Feedback, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Feedback) GetById(id int) (v *Feedback, err error) {
	o := orm.NewOrm()
	v = &Feedback{}
	err = o.QueryTable(new(Feedback)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Feedback) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Feedback) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Feedback) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Feedback) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Feedback) BatchAdd(data []*Feedback) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
