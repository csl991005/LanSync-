package main

import (
	"LanSync/server"
	"os"
	"os/exec"
	"os/signal"
)

func main() {

	go func() {
		server.Run()
	}()

	// 监听退出信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 先写死路径开启 chrome
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:27149/static/index.html")
	cmd.Start()

	// 等待退出信号
	<-chSignal

	cmd.Process.Kill()
}
