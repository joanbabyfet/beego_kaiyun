package service

import (
	"errors"
	"kaiyun/dto"
	"kaiyun/models"
	"kaiyun/utils"
	"strconv"

	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/core/logs"
)

type ArticleService struct {
}

// 获取全部列表
func (s *ArticleService) All(query dto.ArticleQuery) []*models.Article {
	entity := new(models.Article) //new实例化
	return entity.All(query)
}

// 获取分页列表
func (s *ArticleService) PageList(query dto.ArticleQuery) ([]*models.Article, int64) {
	entity := new(models.Article) //new实例化
	return entity.PageList(query)
}

// 获取详情
func (s *ArticleService) GetById(id int) (*models.Article, error) {
	entity := new(models.Article)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("文章不存在 "+strconv.Itoa(id), err)
		return nil, errors.New("文章不存在")
	}
	return info, nil
}

// 保存
func (s *ArticleService) Save(data models.Article) (int, error) {
	stat := 1
	o := orm.NewOrm()
	o.Begin() //开启事务

	if data.Id > 0 {
		//检测数据是否存在
		entity := new(models.Article)
		info, err := entity.GetById(data.Id)
		if err != nil {
			logs.Error("文章不存在 "+strconv.Itoa(data.Id), err)
			return -2, errors.New("文章不存在")
		}
		info.Catid = data.Catid
		info.Title = data.Title
		info.Info = data.Info
		info.Content = data.Content
		info.Author = data.Author
		info.Status = data.Status
		info.UpdateUser = "1"               //修改人
		info.UpdateTime = utils.Timestamp() //修改时间
		ok, _ := info.UpdateById()
		if ok != 1 {
			o.Rollback() //手动回滚事务
			logs.Error("文章更新 "+strconv.Itoa(data.Id), err)
			return -3, errors.New("文章更新失败")
		}
	} else {
		data.Status = 1
		data.CreateUser = "1"               //添加人
		data.CreateTime = utils.Timestamp() //添加时间
		id, _ := data.Add()
		if id <= 0 {
			o.Rollback() //手动回滚事务
			logs.Error("文章添加失败")
			return -4, errors.New("文章添加失败")
		}
	}
	o.Commit() //手动提交事务
	return stat, nil
}

// 删除
// func (s *ArticleService) DeleteById(id int) (int, error) {
// 	stat := 1

// 	entity := new(models.Article)
// 	ok, _ := entity.DeleteById(id)
// 	if ok != 1 {
// 		logs.Error("文章删除 "+strconv.Itoa(id), err)
// 		return -2, errors.New("文章删除失败")
// 	}
// 	return stat, nil
// }

// 软删除
func (s *ArticleService) DeleteById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Article)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("文章不存在 "+strconv.Itoa(id), err)
		return -2, errors.New("文章不存在")
	}

	info.DeleteUser = "1"               //修改人
	info.DeleteTime = utils.Timestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("文章删除 "+strconv.Itoa(id), err)
		return -3, errors.New("文章删除失败")
	}
	return stat, nil
}

// 启用
func (s *ArticleService) EnableById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Article)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("文章不存在 "+strconv.Itoa(id), err)
		return -2, errors.New("文章不存在")
	}

	info.Status = 1
	info.UpdateUser = "1"               //修改人
	info.UpdateTime = utils.Timestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("文章启用 "+strconv.Itoa(id), err)
		return -3, errors.New("文章启用失败")
	}
	return stat, nil
}

// 禁用
func (s *ArticleService) DisableById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Article)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("文章不存在 "+strconv.Itoa(id), err)
		return -2, errors.New("文章不存在")
	}

	info.Status = 0
	info.UpdateUser = "1"               //修改人
	info.UpdateTime = utils.Timestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("文章禁用 "+strconv.Itoa(id), err)
		return -3, errors.New("文章禁用失败")
	}
	return stat, nil
}
