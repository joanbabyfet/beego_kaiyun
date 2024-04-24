package service

import (
	"errors"
	"kaiyun/models"

	"github.com/beego/beego/v2/core/logs"
)

type ContentService struct {
}

// 根据编码获取详情
func (s *ContentService) GetByCode(code string) (*models.Content, error) {
	entity := new(models.Content)
	info, err := entity.GetByCode(code)
	if err != nil {
		logs.Error("内容不存在 "+code, err)
		return nil, errors.New("内容不存在")
	}
	return info, nil
}
