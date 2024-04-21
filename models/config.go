package models

import (
	"kaiyun/dto"
	"reflect"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type Config struct {
	Id          int    `orm:"pk;auto;default();description(ID)" json:"id"`
	Key         string `orm:"size(50);default();null;index;description(键)" json:"key"`
	Name        string `orm:"size(50);default();null;index;description(名称)" json:"name"`
	Value       string `orm:"default();null;description(值)" json:"value"`
	Group       string `orm:"size(50);default();null;description(分组)" json:"group"`
	Type        int    `orm:"default(0);null;description(类型)" json:"type"`
	Description string `orm:"default();null;description(描述)" json:"description"`
	CreateTime  int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser  string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime  int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser  string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime  int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser  string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(Config))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *Config) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "config"
}

// 获取配置文件信息
func (m *Config) GetConfigs(group string) map[string]interface{} {
	list := m.All()
	var mp = make(map[string]interface{})
	for _, config := range list {
		if group == config.Group {
			mp[config.Key] = config.Value
		}
	}
	return mp
}

// 获取全部列表
func (m *Config) All() (list []*Config) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Config))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *Config) PageList(query dto.ConfigQuery) ([]*Config, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Config))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*Config
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*Config, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *Config) GetById(id int) (v *Config, err error) {
	o := orm.NewOrm()
	v = &Config{}
	err = o.QueryTable(new(Config)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *Config) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *Config) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *Config) DeleteById(id int) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *Config) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *Config) BatchAdd(data []*Config) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
