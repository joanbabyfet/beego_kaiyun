package utils

import (
	"github.com/robfig/cron/v3"
)

// 允许往正在执行cron中添加任务, 时间格式 秒 分 时 日 月 周
func CrontabInit() {
	//创建一个定时任务对象 秒级
	crontab := cron.New(cron.WithSeconds())
	crontab.AddFunc("*/10 * * * * *", task1)
	crontab.AddFunc("*/20 * * * * *", task2)
	//启动定时任务 (每个任务会在自己goroutine中执行)
	crontab.Start()
	defer crontab.Stop()
	select {} //查询语句, 阻塞 让main函数不退出, 保持程序运行
}

func task1() {
	//业务
}

func task2() {
	//业务
}
