// beego自带定时任务
package utils

import (
	"github.com/beego/beego/toolbox"
)

// 允许往正在执行cron中添加任务, 时间格式 秒 分 时 日 月 周
func TbCrontabInit() {
	//创建一个定时任务对象 秒级
	tk := toolbox.NewTask("task1", "*/10 * * * * *", TbTask1)
	tk2 := toolbox.NewTask("task2", "*/20 * * * * *", TbTask2)
	toolbox.AddTask("task1", tk)
	toolbox.AddTask("task2", tk2)
	toolbox.StartTask()
	defer toolbox.StopTask()
	select {} //查询语句, 阻塞 让main函数不退出, 保持程序运行
}

func TbTask1() error {
	//业务
	return nil
}

func TbTask2() error {
	//业务
	return nil
}
