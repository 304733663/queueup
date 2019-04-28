package main

import (
	"queueup/libs/queue"
	"queueup/libs/server"
)

func main() {
	go queue.New().Run2()
	server.New().Run()

}
