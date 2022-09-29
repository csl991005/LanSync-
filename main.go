package main

import (
	"LanSync/config"
	"LanSync/server"
	"os"
)

func main() {
	status := make(chan struct{})
	close := make(chan struct{})
	go server.Run()
	go config.OpenChrome(status, close)
	chSignal := config.ListenToInterrupt()
	// 等待退出信号
	for {
		select {
		case <-chSignal:
			close <- struct{}{}
		case <-status:
			os.Exit(0)
		}
	}
}
