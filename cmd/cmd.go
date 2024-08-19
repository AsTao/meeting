package cmd

import (
	"fmt"

	"github.com/AsTao/meeting/conf"
	"github.com/AsTao/meeting/router"
)

func Start() {
	// 初始化配置文件
	conf.InitConfig()
	// 初始化日志
	conf.InitLogger()
	// 初始化数据库
	conf.InitDB()
	// 初始化路由
	router.InitRouter()

}

func Clearn() {
	fmt.Println("=========Clean========")
}
