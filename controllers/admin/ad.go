package admin

import (
	dto "kaiyun/dto"
	"kaiyun/models"
	"kaiyun/service"

	"github.com/beego/beego/v2/core/validation"
)

type AdController struct {
	AdminBaseController
}

// 获取列表
func (c *AdController) Index() {
	catid, _ := c.GetInt("catid")
	page, _ := c.GetInt("page")
	page_size, _ := c.GetInt("page_size")
	if page < 1 {
		page = 1
	}
	if page_size < 1 {
		page_size = 10
	}

	//获取广告列表
	query := dto.AdQuery{}
	query.Catid = catid
	query.Page = page
	query.PageSize = page_size
	query.Status = 1
	service_ad := new(service.AdService)
	list, total := service_ad.PageList(query)

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	var next_page int                    //是否有下一页
	if page*page_size >= int(total) {
		next_page = 0
	} else {
		next_page = 1
	}
	resp["next_page"] = next_page
	resp["total"] = total
	resp["list"] = list
	c.SuccessJson("success", resp)
}

// 获取详情
func (c *AdController) Detail() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Ad{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_ad := new(service.AdService)
	info, err := service_ad.GetById(id)
	if err != nil {
		c.ErrorJson(-2, err.Error(), nil)
	}
	c.SuccessJson("success", info)
}

// 保存
func (c *AdController) Save() {
	id, _ := c.GetInt("id")
	catid, _ := c.GetInt("catid")
	title := c.GetString("title")
	status, _ := c.GetInt8("status")

	//参数验证
	data := models.Ad{
		Id:     id,
		Catid:  catid,
		Title:  title,
		Status: status,
	}
	valid := validation.Validation{}
	if data.Id > 0 {
		valid.Required(data.Id, "id")
	}
	valid.Required(data.Title, "title")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_ad := new(service.AdService)
	stat, err := service_ad.Save(data)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 删除
func (c *AdController) Delete() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Ad{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_ad := new(service.AdService)
	stat, err := service_ad.DeleteById(id)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 启用
func (c *AdController) Enable() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Ad{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_ad := new(service.AdService)
	stat, err := service_ad.EnableById(id)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 禁用
func (c *AdController) Disable() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Ad{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_ad := new(service.AdService)
	stat, err := service_ad.DisableById(id)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}
