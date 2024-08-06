package main

import "github.com/AsTao/meeting.git/cmd"

func main() {

	defer cmd.Clearn()
	cmd.Start()

}
