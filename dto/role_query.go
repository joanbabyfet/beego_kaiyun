package dto

import "kaiyun/utils"

type RoleQuery struct {
	utils.Pager
	Name string `json:"name"`
}
