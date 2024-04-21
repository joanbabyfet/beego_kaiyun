package controllers

import (
	"encoding/json"
	"fmt"
	"kaiyun/dto"
	"kaiyun/models"
	"kaiyun/service"
	"kaiyun/utils"
	"strconv"
	"strings"
)

type MainController struct {
	BaseController
}

// Prepare 初始化数据
func (c *MainController) Prepare() {
	//友情链接
	q := dto.AdQuery{}
	q.Type = 1
	q.Status = 1 //启用
	service_ad := new(service.AdService)
	ads := service_ad.All(q)
	c.Data["ads"] = ads

	//获取分类
	cq := dto.CategoryQuery{}
	cq.Type = 1
	cq.Status = 1 //启用
	service_cat := new(service.CategoryService)
	cats := service_cat.All(cq)
	c.Data["ad_cats"] = cats

	//获取配置
	service_config := new(service.ConfigService)
	mp := service_config.GetConfigs("site")
	mp["logo"] = utils.DisplayImg(mp["logo"].(string))
	mp["orcode_wx"] = utils.DisplayImg(mp["orcode_wx"].(string))
	mp["orcode_wb"] = utils.DisplayImg(mp["orcode_wb"].(string))
	c.Data["config"] = mp

	//获取菜单
	cq = dto.CategoryQuery{}
	cq.Pid = 0
	cq.Type = 0
	cq.Status = 1 //启用
	service_cat = new(service.CategoryService)
	pCate := service_cat.All(cq)
	c.Data["pCate"] = pCate
}

func (c *MainController) Index() {
	//获取首页分类
	service_cat := new(service.CategoryService)
	cat, _ := service_cat.GetById(1)
	imgs := strings.Split(cat.Banner, ",")
	var imgMap []map[string]string
	for _, img := range imgs {
		cMap := map[string]string{}
		cMap["url"] = utils.DisplayImg(img)
		imgMap = append(imgMap, cMap)
	}
	c.Data["imgs"] = imgMap

	//获取关于我们内容
	code := "about"
	service_content := new(service.ContentService)
	about, _ := service_content.GetByCode(code)
	about.Video = utils.DisplayVideo(about.Video)
	c.Data["about"] = about

	//获取产品
	pq := dto.ProductQuery{}
	pq.Catid = 16 //技术服务
	pq.Status = 1 //启用
	service_product := new(service.ProductService)
	products := service_product.All(pq)
	if len(products) > 0 {
		map_content := make(map[string]interface{})
		err := json.Unmarshal([]byte(products[0].Content), &map_content)
		if err != nil {
			return
		}
		c.Data["service"] = map_content
	}

	//获取分类ids
	cq := dto.CategoryQuery{}
	cq.Pid = 19   //自主产品
	cq.Status = 1 //启用
	service_cat = new(service.CategoryService)
	cats := service_cat.All(cq)
	var ids []int
	for _, cat := range cats {
		ids = append(ids, cat.Id)
	}

	//获取自主产品
	pq = dto.ProductQuery{}
	pq.Catids = ids
	pq.Status = 1 //启用
	service_product = new(service.ProductService)
	products = service_product.All(pq)
	//数据格式化
	var product_list []*models.Product
	for _, product := range products {
		product.Img = utils.DisplayImg(product.Img)
		product_list = append(product_list, product)
	}
	c.Data["products"] = product_list

	//获取合作伙伴
	pq = dto.ProductQuery{}
	pq.Limit = 12
	pq.Catid = 22
	pq.Status = 1 //启用
	service_product = new(service.ProductService)
	products = service_product.All(pq)
	var peers []map[string]string
	for _, product := range products {
		var p []map[string]string
		json.Unmarshal([]byte(product.Content), &p)
		peers = append(peers, p...)
	}
	c.Data["partners"] = peers

	//获取重要案例
	aq := dto.ArticleQuery{}
	aq.Catid = 11
	aq.Limit = 2
	aq.Status = 1 //启用
	service_article := new(service.ArticleService)
	articles := service_article.All(aq)
	//数据格式化
	var article_list []*models.Article
	for _, article := range articles {
		article.Img = utils.DisplayImg(article.Img)
		article_list = append(article_list, article)
	}
	c.Data["articles"] = article_list

	c.Layout = "home_layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "index.html"
}

func (c *MainController) About() {
	//获取关于我们内容
	code := "about"
	service_content := new(service.ContentService)
	about, _ := service_content.GetByCode(code)
	about.Video = utils.DisplayVideo(about.Video)
	c.Data["about"] = about
	var bs []map[string]string
	json.Unmarshal([]byte(about.Bs), &bs) //json字符串转struct
	c.Data["bs"] = bs

	//获取主营业务
	aq := dto.ArticleQuery{}
	aq.Catid = 16
	aq.Status = 1 //启用
	service_article := new(service.ArticleService)
	articles := service_article.All(aq)
	//数据格式化
	var articles_plus []*models.Article
	for _, article := range articles {
		article.Img = utils.DisplayImg(article.Img)
		articles_plus = append(articles_plus, article)
	}
	c.Data["articles"] = articles_plus

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "about.html"
}

func (c *MainController) Case() {
	page, _ := c.GetInt("page")
	if page < 1 {
		page = 1
	}

	//行业客户
	pq := dto.ProductQuery{}
	pq.Catid = 12
	pq.Status = 1 //启用
	service_product := new(service.ProductService)
	products := service_product.All(pq)
	var imgs []map[string]string
	if len(products) > 0 {
		err := json.Unmarshal([]byte(products[0].Content), &imgs)
		if err != nil {
			return
		}
		c.Data["imgs"] = imgs
	}

	//获取主要案例分页数据
	aq := dto.ArticleQuery{}
	aq.Page = page
	aq.PageSize = 4
	aq.Catid = 11
	aq.Status = 1 //启用
	service_article := new(service.ArticleService)
	articles, count := service_article.PageList(aq)
	//数据格式化
	var articles_plus []*models.Article
	for _, article := range articles {
		article.Img = utils.DisplayImg(article.Img)
		articles_plus = append(articles_plus, article)
	}
	c.Data["articles"] = articles_plus
	//分页器
	c.Data["pagebar"] = utils.NewPager(page, int(count), aq.PageSize, "/case", true).ToString()

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "case.html"
}

func (c *MainController) Contact() {
	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "contact.html"
}

// 内容详情
func (c *MainController) Content() {
	id, _ := strconv.Atoi(c.GetString(":id"))
	index, _ := c.GetInt("index")
	if index == 0 {
		service_article := new(service.ArticleService)
		one, _ := service_article.GetById(id)
		c.Data["one"] = one
	} else {
		service_product := new(service.ProductService)
		one, _ := service_product.GetById(id)
		c.Data["one"] = one
	}

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "content.html"
}

func (c *MainController) Search() {
	keyword := c.GetString("keyword")
	page, _ := c.GetInt("page")
	index, _ := c.GetInt("index")
	page_size, _ := c.GetInt("page_size")
	if page < 1 {
		page = 1
	}
	if page_size < 1 {
		page_size = 10
	}
	c.Data["keyword"] = keyword
	c.Data["index"] = index

	var count int
	if index == 0 { //默认搜客户和案例
		aq := dto.ArticleQuery{}
		aq.Title = keyword
		aq.Page = page
		aq.PageSize = page_size
		aq.Status = 1
		service_article := new(service.ArticleService)
		articles, total := service_article.PageList(aq)
		//数据格式化
		var list []*models.Article
		for _, article := range articles {
			article.Img = utils.DisplayImg(article.Img)
			list = append(list, article)
		}
		c.Data["data"] = list
		count = int(total)
	} else { //自主产品
		pq := dto.ProductQuery{}
		pq.Title = keyword
		pq.Page = page
		pq.PageSize = page_size
		pq.Status = 1
		service_product := new(service.ProductService)
		products, total := service_product.PageList(pq)
		//数据格式化
		var list []*models.Product
		for _, product := range products {
			product.Img = utils.DisplayImg(product.Img)
			list = append(list, product)
		}
		c.Data["data"] = list
		count = int(total)
	}
	//总条数
	c.Data["count"] = count
	//分页器
	c.Data["pagebar"] = utils.NewPager(page, count, page_size,
		fmt.Sprintf("/search?keyword=%s&index=%d", keyword, index), true).ToString()

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "search.html"
}

func (c *MainController) Partner() {
	//获取产品
	pq := dto.ProductQuery{}
	pq.Catid = 22
	pq.Status = 1 //启用
	service_product := new(service.ProductService)
	products := service_product.All(pq)
	var mp []map[string]interface{}
	for _, product := range products {
		p := make(map[string]interface{})
		p["Title"] = product.Title
		var imgs []map[string]string
		json.Unmarshal([]byte(product.Content), &imgs)
		p["Imgs"] = imgs
		mp = append(mp, p)
	}
	c.Data["partners"] = mp

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "partner.html"
}

func (c *MainController) Service() {
	//获取产品
	pq := dto.ProductQuery{}
	pq.Catid = 16 //技术服务
	pq.Status = 1 //启用
	service_product := new(service.ProductService)
	products := service_product.All(pq)
	if len(products) > 0 {
		map_content := make(map[string]interface{})
		err := json.Unmarshal([]byte(products[0].Content), &map_content)
		if err != nil {
			return
		}
		c.Data["service"] = map_content
	}

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "service.html"
}

func (c *MainController) Product() {
	//获取分类ids
	cq := dto.CategoryQuery{}
	cq.Pid = 19   //自主产品
	cq.Status = 1 //启用
	service_cat := new(service.CategoryService)
	cats := service_cat.All(cq)
	var ids []int
	for _, cat := range cats {
		ids = append(ids, cat.Id)
	}

	//获取产品
	pq := dto.ProductQuery{}
	pq.Catids = ids
	pq.Status = 1 //启用
	service_product := new(service.ProductService)
	products := service_product.All(pq)
	//数据格式化
	var list []*models.Product
	for _, product := range products {
		product.Img = utils.DisplayImg(product.Img)
		list = append(list, product)
	}
	c.Data["products"] = list
	c.Data["Title"] = "自主产品"

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "product.html"
}

func (c *MainController) Quality() {
	cid, _ := strconv.Atoi(c.GetString("cid"))
	c.Data["currentId"] = cid //当前分类id

	//获取分类ids
	cq := dto.CategoryQuery{}
	cq.Pid = 6    //公司资质
	cq.Status = 1 //启用
	service_cat := new(service.CategoryService)
	cats := service_cat.All(cq)
	var ids []int
	var sonCates []*models.Category
	for _, cat := range cats {
		ids = append(ids, cat.Id)
		sonCates = append(sonCates, cat)
	}
	c.Data["qcate"] = sonCates

	//获取公司资质
	aq := dto.ArticleQuery{}
	aq.Catids = ids
	aq.Status = 1 //启用
	service_article := new(service.ArticleService)
	articles := service_article.All(aq)
	//数据格式化
	var list []*models.Article
	for _, article := range articles {
		article.Img = utils.DisplayImg(article.Img)
		list = append(list, article)
	}
	c.Data["articles"] = list

	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.TplName = "quality.html"
}
