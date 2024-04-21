// 分页器
package utils

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Pager struct {
	Page          int   `json:"page"`
	PageSize      int   `json:"pageSize"`
	TotalElements int64 `json:"totalElements"`
	TotalPages    int   `json:"totalPages"`
	urlpath       string
	urlquery      string
	nopath        bool
}

func NewPager(page, totalnum, pagesize int, url string, nopath ...bool) *Pager {
	p := new(Pager)
	p.Page = page
	p.TotalElements = int64(totalnum)
	p.PageSize = pagesize

	arr := strings.Split(url, "?")
	p.urlpath = arr[0]
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	} else {
		p.urlquery = ""
	}

	if len(nopath) > 0 {
		p.nopath = nopath[0]
	} else {
		p.nopath = false
	}

	return p
}

func (u *Pager) url(page int) string {
	if u.nopath { //不使用目录形式
		if u.urlquery != "" {
			return fmt.Sprintf("%s%s&page=%d", u.urlpath, u.urlquery, page)
		} else {
			return fmt.Sprintf("%s?page=%d", u.urlpath, page)
		}
	} else {
		return fmt.Sprintf("%s/page/%d%s", u.urlpath, page, u.urlquery)
	}
}

func (u *Pager) ToString() string {
	if u.TotalElements <= int64(u.PageSize) {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int

	offset = 5
	linknum = 10

	totalpage = int(math.Ceil(float64(u.TotalElements) / float64(u.PageSize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = u.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	if u.Page > 1 {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-prev\" href=\"%s\"><</a></li>", u.url(u.Page-1)))
	} else {
		buf.WriteString("<span><</span>")
	}

	if u.Page > linknum {
		buf.WriteString(fmt.Sprintf("<a href=\"%s\" class=\"laypage_first\">1...</a>", u.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == u.Page {
			buf.WriteString(fmt.Sprintf("<span class=\"layui-laypage-curr\"><em class=\"layui-laypage-em\"></em><em>%d</em></span>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<a href=\"%s\">%d</a>", u.url(i), i))
		}
	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-last\" href=\"%s\">末页</a>", u.url(totalpage)))
	}

	if u.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<a class=\"layui-laypage-next\" href=\"%s\">></a></li>", u.url(u.Page+1)))
	} else {
		buf.WriteString("<span>></span>")
	}
	return buf.String()
}
