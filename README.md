## About
使用 beego v2 搭建kaiyun官网前台, 支持响应式, 提供关于我们/客户和案例/技术服务/产品列表/合作伙伴/联系我们等

## Feature
* JWT身份验证
* 文件上传与缩略图功能
* 提供图片验证码功能
* 发送短信功能
* 发送邮件功能
* 发送chatGPT功能
* 日志功能
* 实现跨域解决
* i18n多语言功能
* crontab定时任务功能
* WebSocket服务端

## Usage
```
# 执行
bee run
# windows 打包
bee pack -be GOOS=windows
# linux 打包
bee pack -be GOOS=linux
```

## Requires
go v1.19.3
bee v2.1.0

## Maintainers
Alan

## LICENSE
[MIT License](https://github.com/joanbabyfet/beego_kaiyun/blob/master/LICENSE)
