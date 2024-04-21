package models

import (
	"github.com/beego/beego"
	"github.com/beego/beego/orm"

	"kaiyun/dto"
	"reflect"
)

// 定义结构体, 字段首字母要大写才能进行json解析, 会自动转蛇底命令例 create_user
type User struct {
	Id           string `orm:"pk;size(32);default();description(ID)" json:"id"`
	Origin       int8   `orm:"default(0);null;description(注册来源 1=H5 2=PC)" json:"origin"`
	Username     string `orm:"size(40);default();null;index;description(帐号)" json:"username"`
	Password     string `orm:"size(60);default();null;description(密码)" json:"_"`
	Avatar       string `orm:"size(100);default();null;description(头像)" json:"avatar"`
	Realname     string `orm:"size(50);default();null;index;description(姓名)" json:"realname"`
	Sex          int8   `orm:"default(1);null;description(性别 0=女 1=男)" json:"sex"`
	Email        string `orm:"size(100);default();null;index;description(信箱)" json:"email"`
	PhoneCode    string `orm:"size(5);default();null;index;description(手机号国码)" json:"phone_code"`
	Phone        string `orm:"size(20);default();null;index;description(手机号)" json:"phone"`
	Address      string `orm:"size(100);default();null;description(地址)" json:"address"`
	Salt         string `orm:"size(128);default();null;description(加密钥匙)" json:"salt"`
	RoleId       int    `orm:"default(0);null;description(角色)" json:"role_id"`
	RegIp        string `orm:"size(15);default();null;description(注册ip)" json:"reg_ip"`
	LoginTime    int    `orm:"default(0);null;description(最后登录时间)" json:"login_time"`
	LoginIp      string `orm:"size(15);default();null;description(最后登录IP)" json:"login_ip"`
	LoginCountry string `orm:"size(2);default();null;description(最后登录国家)" json:"login_country"`
	Language     string `orm:"size(10);default();null;description(语言)" json:"language"`
	Status       int8   `orm:"default(1);null;description(状态: 0=禁用 1=启用)" json:"status"`
	CreateTime   int    `orm:"default(0);null;description(創建時間)" json:"create_time"`
	CreateUser   string `orm:"size(32);default(0);null;description(創建人)" json:"create_user"`
	UpdateTime   int    `orm:"default(0);null;description(修改時間)" json:"update_time"`
	UpdateUser   string `orm:"size(32);default(0);null;description(修改人)" json:"update_user"`
	DeleteTime   int    `orm:"default(0);null;description(刪除時間)" json:"delete_time"`
	DeleteUser   string `orm:"size(32);default(0);null;description(刪除人)" json:"delete_user"`
}

// 修改密码请求格式
type EditPwdReq struct {
	Id          string `json:"id"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
}

// 登录请求格式
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Captcha  string `json:"captcha"`
}

// 登录响应格式
type LoginResp struct {
	Id        string `json:"id"`
	Origin    int8   `json:"origin"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Realname  string `json:"realname"`
	Email     string `json:"email"`
	PhoneCode string `json:"phone_code"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	//Salt         string `json:"salt"`
	RegIp        string `json:"reg_ip"`
	LoginTime    int64  `json:"login_time"`
	LoginIp      string `json:"login_ip"`
	LoginCountry string `json:"login_country"`
	Language     string `json:"language"`
	Status       int8   `json:"status"`
	CreateTime   int64  `json:"create_time"`
	CreateUser   string `json:"create_user"`
	UpdateTime   int64  `json:"update_time"`
	UpdateUser   string `json:"update_user"`
	DeleteTime   int64  `json:"delete_time"`
	DeleteUser   string `json:"delete_user"`
	AccessToken  string `json:"access_token"`
}

// 在models里注册模型
func init() {
	orm.RegisterModel(new(User))
}

// 重写TableName方法，返回对应数据库中的表名
func (m *User) TableName() string {
	db_prefix := beego.AppConfig.String("db_prefix")
	return db_prefix + "user"
}

// 获取全部列表
func (m *User) All(query dto.UserQuery) (list []*User) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	qs = qs.Filter("delete_time", 0) //未删除
	_, err := qs.OrderBy("-create_time").All(&list)
	if err != nil {
		return nil
	}
	return list
}

// 获取分页列表
func (m *User) PageList(query dto.UserQuery) ([]*User, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	qs = qs.Filter("delete_time", 0) //未删除
	//总条数
	count, _ := qs.Count()
	var list []*User
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-create_time").Limit(query.PageSize, offset).All(&list)
	}
	if reflect.ValueOf(list).IsNil() {
		list = make([]*User, 0) //赋值为空切片[]
	}
	return list, count
}

// 获取单条
func (m *User) GetById(id string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{}
	err = o.QueryTable(new(User)).Filter("delete_time", 0).Filter("id", id).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 单条添加
func (m *User) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// 更新
func (m *User) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// 删除
func (m *User) DeleteById(id string) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// 添加或更新
func (m *User) InsertOrUpdate() (int64, error) {
	o := orm.NewOrm()
	return o.InsertOrUpdate(m)
}

// 批量添加 (支持多条插入数据库 例 mysql)
func (m *User) BatchAdd(data []*User) (int64, error) {
	o := orm.NewOrm()
	return o.InsertMulti(len(data), data)
}
