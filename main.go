package main

import (
	"github.com/AsTao/meeting/cmd"
)

func main() {

	defer cmd.Clearn()
	cmd.Start()

}
