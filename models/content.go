package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Content struct {
	Id         int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Code       string `orm:"size(20);default();null;description(编码)" json:"code"`
	Title      string `orm:"size(50);default();null;index;description(标题)" json:"title"`
	Img        string `orm:"size(100);default();null;description(banner图片)" json:"img"`
	TopImg     string `orm:"size(100);default();null;description(导航图片)" json:"top_img"`
	Video      string `orm:"size(100);default();null;description(视频地址)" json:"video"`
	Content    string `orm:"type(text);null;description(內容)" json:"content"`
	Bs         string `orm:"type(text);null;description(中部业务)" json:"bs"`
	CreateTime int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(Content))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Content) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "content"
}

// 获取全部列表
func (m *Content) All(query dto.ContentQuery) (list []*Content) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Content))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Content) PageList(query dto.ContentQuery) ([]*Content, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Content))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*Content
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Content, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Content) GetById(id int) (v *Content, err error) {
	o := orm.NewOrm()
	v = &Content{}
	err = o.QueryTable(new(Content)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 根据编码获取单条
func (m *Content) GetByCode(code string) (v *Content, err error) {
	o := orm.NewOrm()
	v = &Content{}
	err = o.QueryTable(new(Content)).Filter("delete_time", 0).Filter("code", code).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Content) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Content) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Content) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Content) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Content) BatchAdd(data []*Content) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
