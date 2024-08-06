package cmd

import (
	"fmt"

	"github.com/AsTao/meeting.git/conf"
)

func Start() {
	conf.InitConfig()
}

func Clearn() {
	fmt.Println("=========Clean========")
}
