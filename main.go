package main

import (
	"os"

	"github.com/weilaim/blog-api/conf"
	"github.com/weilaim/blog-api/server"
	"github.com/weilaim/blog-api/tasks"
)

func main(){
	//读取配置
	conf.Init()

	//定时任务
	tasks.CronJob()

	//装载路由
	r := server.NewRouter()
	r.Run(":"+os.Getenv("SERVER_PORT"))
}