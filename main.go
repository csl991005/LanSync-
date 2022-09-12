package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		r := gin.Default()
		r.GET("/", func(ctx *gin.Context) {
			ctx.Writer.Write([]byte("test message"))
		})
		r.Run(":8080")
	}()

	// 监听退出信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 先写死路径开启 chrome
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/")
	cmd.Start()

	// 等待退出信号
	<-chSignal

	cmd.Process.Kill()
}
