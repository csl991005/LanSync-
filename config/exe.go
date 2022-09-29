package config

import (
	"os"
	"os/exec"
	"os/signal"
)

func OpenChrome(status chan struct{}, close chan struct{}) {
	// 先写死路径开启 chrome
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:27149/static/index.html")
	cmd.Start()
	go func() {
		<-close
		cmd.Process.Kill()
	}()
	cmd.Wait()
	status <- struct{}{}
}

func ListenToInterrupt() chan os.Signal {
	// 监听退出信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
