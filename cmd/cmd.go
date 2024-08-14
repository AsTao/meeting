package cmd

import (
	"fmt"

	"github.com/AsTao/meeting/conf"
	"github.com/AsTao/meeting/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clearn() {
	fmt.Println("=========Clean========")
}
