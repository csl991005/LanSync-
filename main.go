package main

import (
	"LanSync/config"
	"LanSync/server"
	"os"
)

func main() {
	status := make(chan struct{})
	go server.Run()
	go config.OpenChrome(status)
	chSignal := config.ListenToInterrupt()
	// 等待退出信号
	select {
	case <-chSignal:
	case <-status:
		os.Exit(1)
	}

}
