package cmd

import (
	"fmt"

	"github.com/AsTao/meeting.git/conf"
	"github.com/AsTao/meeting.git/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clearn() {
	fmt.Println("=========Clean========")
}
