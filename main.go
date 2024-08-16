package main

import (
	"github.com/AsTao/meeting/cmd"
)

// @title 会议室预定 tao
// @version 0.0.1
// @description swagger api
func main() {

	defer cmd.Clearn()
	cmd.Start()

}
