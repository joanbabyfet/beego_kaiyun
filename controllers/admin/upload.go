package admin

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/beego/beego"
	"github.com/disintegration/imaging"
)

type UploadController struct {
	AdminBaseController
}

// 上传图片
func (c *UploadController) Upload() {
	f, h, err := c.GetFile("filename") //获取文件信息
	dir := c.GetString("dir")          //文件上传目录, 默认image
	thumb_w, _ := c.GetInt("thumb_w")  //缩图宽
	thumb_h, _ := c.GetInt("thumb_h")  //缩图高
	if err != nil {
		c.ErrorJson(-1, "上传文件失败", nil)
	}
	if dir == "" {
		dir = "image"
	}
	file_url := beego.AppConfig.String("file_url")

	//文件大小校验
	upload_max_size := beego.AppConfig.String("upload_max_size")
	max_size, _ := strconv.ParseInt(upload_max_size, 10, 64)
	if h.Size > max_size*1024*1024 {
		c.ErrorJson(-2, "您上传的文件过大,最大值为"+upload_max_size+"MB", nil)
	}

	//文件后缀过滤
	ext := path.Ext(h.Filename) //输出.jpg
	allow_ext_map := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := allow_ext_map[ext]; !ok {
		c.ErrorJson(-3, "文件格式不正确", nil)
	}

	//创建目录
	upload_dir := beego.AppConfig.String("upload_dir")
	dir_num := time.Now().Format("20060102") //输出 20240404/
	err = os.MkdirAll(upload_dir+"/"+dir+"/"+dir_num, os.FileMode(0775))
	if err != nil {
		c.ErrorJson(-4, "创建目录失败", nil)
	}

	//构造文件名
	source := rand.NewSource(time.Now().UnixNano()) //这里用系统时间毫秒值当种子值
	r := rand.New(source)
	rand_num := fmt.Sprintf("%d", r.Intn(9999)+1000) //获取1000-9999随机数
	hash_name := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + rand_num))
	file_name := fmt.Sprintf("%x", hash_name) + ext //文件名 例 cf386af3f37962ad3769054f68d7a049.jpg

	path := upload_dir + "/" + dir + "/" + dir_num + "/" + file_name
	filelink := file_url + "/" + dir + "/" + dir_num + "/" + file_name
	defer f.Close()                // 延迟关闭文件流, 否则会出现临时文件不能清除的情况
	c.SaveToFile("filename", path) //名称与 c.GetFile("xxx") 一致

	//生成缩略图
	if thumb_w > 0 || thumb_h > 0 {
		src, err := imaging.Open(path)
		if err != nil {
			c.ErrorJson(-5, "开启缩略图失败", nil)
		}
		dsc := imaging.Resize(src, thumb_w, thumb_h, imaging.Lanczos)

		//构造缩略图文件名
		source := rand.NewSource(time.Now().UnixNano()) //这里用系统时间毫秒值当种子值
		r := rand.New(source)
		rand_num := fmt.Sprintf("%d", r.Intn(9999)+1000) //获取1000-9999随机数
		hash_name := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + rand_num))
		file_name = fmt.Sprintf("%x", hash_name) + ext //文件名 例 cf386af3f37962ad3769054f68d7a049.jpg
		path = upload_dir + "/" + dir + "/" + dir_num + "/" + file_name
		filelink = file_url + "/" + dir + "/" + dir_num + "/" + file_name

		err = imaging.Save(dsc, path)
		if err != nil {
			c.ErrorJson(-6, "生成缩略图失败", nil)
		}
	}

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["realname"] = h.Filename
	resp["filename"] = dir_num + "/" + file_name
	resp["filelink"] = filelink
	c.SuccessJson("success", resp) //返回空对象
}

// 下载文件
func (c *UploadController) Download() {
	//获取文件名 20240404/beaedec7c974a5c8e9a9f8770f9cec2b.png
	filename := "20240420/d28e41300d966ce4693e8d3236bc9552.jpg"

	c.Ctx.Output.Download("uploads/", filename)
}
