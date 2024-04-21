package service

import "kaiyun/models"

type ConfigService struct {
}

// 获取配置文件信息
func (s *ConfigService) GetConfigs(group string) map[string]interface{} {
	entity := new(models.Config) //new实例化
	return entity.GetConfigs(group)
}
