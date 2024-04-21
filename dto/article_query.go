package dto

import "kaiyun/utils"

type ArticleQuery struct {
	utils.Pager
	Catid  int    `json:"catid"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Limit  int    `json:"limit"`
	Catids []int  `json:"catids"`
}
