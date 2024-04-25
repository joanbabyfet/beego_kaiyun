// 公共方法
package utils

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/beego/beego"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/mojocn/base64Captcha"
	"github.com/sashabaranov/go-openai"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

// 验证码存放
var Store = base64Captcha.DefaultMemStore

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	//占位待%x为整型以十六进制方式显示
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// 密码加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// 密码验证
func PasswordVerify(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

// 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

// 获取时间戳
func Timestamp() int {
	t := time.Now().Unix()
	return int(t)
}

// 获取当前日期时间
func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前日期
func Date() string {
	return time.Now().Format("2006-01-02")
}

// 时间戳转日期
func UnixToDateTime(timestramp int) string {
	t := time.Unix(int64(timestramp), 0)
	return t.Format("2006-01-02 15:04:05") //通用时间模板定义
}

// 时间戳转日期
func UnixToDate(timestramp int) string {
	t := time.Unix(int64(timestramp), 0)
	return t.Format("2006-01-02") //通用时间模板定义
}

// 日期转时间戳
func DateToUnix(str string) int {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return 0
	}
	return int(t.Unix())
}

// 发送邮件
func SendMail(to string, subject string, body string) bool {
	mail_host := beego.AppConfig.String("mail_host")
	mail_port, _ := strconv.Atoi(beego.AppConfig.String("mail_port"))
	mail_username := beego.AppConfig.String("mail_username")
	mail_password := beego.AppConfig.String("mail_password")
	mail_from_address := beego.AppConfig.String("mail_from_address")
	mail_from_name := beego.AppConfig.String("mail_from_name")

	m := gomail.NewMessage()
	m.SetAddressHeader("From", mail_from_address, mail_from_name)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mail_host, mail_port, mail_username, mail_password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		logs.Error(err)
		//panic(err) //后续代码不会执行
		return false
	}
	return true
}

// 发送短信
func SendSMS(to string, body string) bool {
	twilio_sid := beego.AppConfig.String("twilio_sid")
	twilio_token := beego.AppConfig.String("twilio_token")
	twilio_from := beego.AppConfig.String("twilio_from")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilio_sid,
		Password: twilio_token,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(twilio_from)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		logs.Error(err)
		return false
	}
	response, _ := json.Marshal(*resp) //将数据编码为json字符串, 返回[]uint8
	//返回json字符串写入日志
	logs.Info(string(response))

	return true
}

// 发送ChatGPT
func ChatGPT(keyword string) (bool, string) {
	api := beego.AppConfig.String("chat_gpt_api")
	client := openai.NewClient(api)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: keyword,
				},
			},
		},
	)
	if err != nil {
		logs.Error(err)
		return false, ""
	}
	result := resp.Choices[0].Message.Content
	return true, result
}

// 设置表单验证信息
func SetVerifyMessage() {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "不能为空",
		"Min":          "最小值 为 %d",
		"Max":          "最大值 为 %d",
		"Range":        "范围 为 %d 到 %d",
		"MinSize":      "最短长度 为 %d",
		"MaxSize":      "最大长度 为 %d",
		"Length":       "长度必须 为 %d",
		"Alpha":        "必须是有效的字母",
		"Numeric":      "必须是有效的数字",
		"AlphaNumeric": "必须是有效的字母或数字",
		"Match":        "必须匹配 %s",
		"NoMatch":      "必须不匹配 %s",
		"AlphaDash":    "必须是有效的字母、数字或连接符号(-_)",
		"Email":        "必须是有效的电子邮件地址",
		"IP":           "必须是有效的IP地址",
		"Base64":       "必须是有效的base64字符",
		"Mobile":       "必须是有效的手机号码",
		"Tel":          "必须是有效的电话号码",
		"Phone":        "必须是有效的电话或移动电话号码",
		"ZipCode":      "必须是有效的邮政编码",
	})
}

// 获取图片完整地址
func DisplayImg(filename string) string {
	if filename == "" {
		return ""
	}
	return beego.AppConfig.String("file_url") + "/image/" + filename
}

// 获取视频完整地址
func DisplayVideo(filename string) string {
	if filename == "" {
		return ""
	}
	return beego.AppConfig.String("file_url") + "/video/" + filename
}

// 获取图片验证码
func GetCaptcha() (string, string, string, error) {
	driver := base64Captcha.DefaultDriverDigit
	captcha := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, answer, err := captcha.Generate()
	return id, b64s, answer, err
}
