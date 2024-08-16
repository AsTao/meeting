package main

import (
	"github.com/AsTao/meeting/cmd"
)

// @title Go-web tao
// @version 0.0.1
// @description 会议预定
func main() {

	defer cmd.Clearn()
	cmd.Start()

}
