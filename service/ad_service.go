package service

import (
	"errors"
	"kaiyun/dto"
	"kaiyun/models"
	"kaiyun/utils"

	"github.com/beego/beego/logs"
)

type AdService struct {
}

// 获取全部列表
func (s *AdService) All(query dto.AdQuery) []*models.Ad {
	entity := new(models.Ad) //new实例化
	return entity.All(query)
}

// 获取分页列表
func (s *AdService) PageList(query dto.AdQuery) ([]*models.Ad, int64) {
	entity := new(models.Ad) //new实例化
	return entity.PageList(query)
}

// 获取详情
func (s *AdService) GetById(id int) (*models.Ad, error) {
	entity := new(models.Ad)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("该数据不存在")
		return nil, errors.New("该数据不存在")
	}
	return info, nil
}

// 保存
func (s *AdService) Save(data models.Ad) (int, error) {
	stat := 1

	if data.Id > 0 {
		//检测数据是否存在
		entity := new(models.Ad)
		info, err := entity.GetById(data.Id)
		if err != nil {
			logs.Error("数据不存在")
			return -2, errors.New("数据不存在")
		}
		info.Catid = data.Catid
		info.Title = data.Title
		info.Status = data.Status
		info.UpdateUser = "1"                  //修改人
		info.UpdateTime = utils.GetTimestamp() //修改时间
		ok, _ := info.UpdateById()
		if ok != 1 {
			logs.Error("更新失败")
			return -3, errors.New("更新失败")
		}
	} else {
		data.Status = 1
		data.CreateUser = "1"                  //添加人
		data.CreateTime = utils.GetTimestamp() //添加时间
		id, _ := data.Add()
		if id <= 0 {
			logs.Error("添加失败")
			return -4, errors.New("添加失败")
		}
	}
	return stat, nil
}

// 删除
// func (s *AdService) DeleteById(id int) (int, error) {
// 	stat := 1

// 	entity := new(models.Ad)
// 	ok, _ := entity.DeleteById(id)
// 	if ok != 1 {
// 		logs.Error("删除失败")
// 		return -2, errors.New("删除失败")
// 	}
// 	return stat, nil
// }

// 软删除
func (s *AdService) DeleteById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Ad)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("数据不存在")
		return -2, errors.New("数据不存在")
	}

	info.DeleteUser = "1"                  //修改人
	info.DeleteTime = utils.GetTimestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("删除失败")
		return -3, errors.New("删除失败")
	}
	return stat, nil
}

// 启用
func (s *AdService) EnableById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Ad)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("数据不存在")
		return -2, errors.New("数据不存在")
	}

	info.Status = 1
	info.UpdateUser = "1"                  //修改人
	info.UpdateTime = utils.GetTimestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("启用失败")
		return -3, errors.New("启用失败")
	}
	return stat, nil
}

// 禁用
func (s *AdService) DisableById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Ad)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("数据不存在")
		return -2, errors.New("数据不存在")
	}

	info.Status = 0
	info.UpdateUser = "1"                  //修改人
	info.UpdateTime = utils.GetTimestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("禁用失败")
		return -3, errors.New("禁用失败")
	}
	return stat, nil
}
