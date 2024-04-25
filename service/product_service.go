package service

import (
	"errors"
	"kaiyun/dto"
	"kaiyun/models"
	"kaiyun/utils"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
)

type ProductService struct {
}

// 获取全部列表
func (s *ProductService) All(query dto.ProductQuery) []*models.Product {
	entity := new(models.Product) //new实例化
	return entity.All(query)
}

// 获取分页列表
func (s *ProductService) PageList(query dto.ProductQuery) ([]*models.Product, int64) {
	entity := new(models.Product) //new实例化
	return entity.PageList(query)
}

// 获取详情
func (s *ProductService) GetById(id int) (*models.Product, error) {
	entity := new(models.Product)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("产品不存在 "+strconv.Itoa(id), err)
		return nil, errors.New("产品不存在")
	}
	return info, nil
}

// 保存
func (s *ProductService) Save(data models.Product) (int, error) {
	stat := 1

	if data.Id > 0 {
		//检测数据是否存在
		entity := new(models.Product)
		info, err := entity.GetById(data.Id)
		if err != nil {
			logs.Error("产品不存在 "+strconv.Itoa(data.Id), err)
			return -2, errors.New("产品不存在")
		}
		info.Catid = data.Catid
		info.Title = data.Title
		info.Status = data.Status
		info.UpdateUser = "1"               //修改人
		info.UpdateTime = utils.Timestamp() //修改时间
		ok, _ := info.UpdateById()
		if ok != 1 {
			logs.Error("产品更新 "+strconv.Itoa(data.Id), err)
			return -3, errors.New("产品更新失败")
		}
	} else {
		data.Status = 1
		data.CreateUser = "1"               //添加人
		data.CreateTime = utils.Timestamp() //添加时间
		id, _ := data.Add()
		if id <= 0 {
			logs.Error("产品添加失败")
			return -4, errors.New("产品添加失败")
		}
	}
	return stat, nil
}

// 删除
// func (s *ProductService) DeleteById(id int) (int, error) {
// 	stat := 1

// 	entity := new(models.Product)
// 	ok, _ := entity.DeleteById(id)
// 	if ok != 1 {
// 		logs.Error("产品删除 "+strconv.Itoa(id), err)
// 		return -2, errors.New("产品删除失败")
// 	}
// 	return stat, nil
// }

// 软删除
func (s *ProductService) DeleteById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Product)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("产品不存在 "+strconv.Itoa(id), err)
		return -2, errors.New("产品不存在")
	}

	info.DeleteUser = "1"               //修改人
	info.DeleteTime = utils.Timestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("产品删除 "+strconv.Itoa(id), err)
		return -3, errors.New("产品删除失败")
	}
	return stat, nil
}

// 启用
func (s *ProductService) EnableById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Product)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("产品不存在 "+strconv.Itoa(id), err)
		return -2, errors.New("产品不存在")
	}

	info.Status = 1
	info.UpdateUser = "1"               //修改人
	info.UpdateTime = utils.Timestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("产品启用 "+strconv.Itoa(id), err)
		return -3, errors.New("产品启用失败")
	}
	return stat, nil
}

// 禁用
func (s *ProductService) DisableById(id int) (int, error) {
	stat := 1

	//检测数据是否存在
	entity := new(models.Product)
	info, err := entity.GetById(id)
	if err != nil {
		logs.Error("产品不存在 "+strconv.Itoa(id), err)
		return -2, errors.New("产品不存在")
	}

	info.Status = 0
	info.UpdateUser = "1"               //修改人
	info.UpdateTime = utils.Timestamp() //修改时间
	ok, _ := info.UpdateById()
	if ok != 1 {
		logs.Error("产品禁用 "+strconv.Itoa(id), err)
		return -3, errors.New("产品禁用失败")
	}
	return stat, nil
}
