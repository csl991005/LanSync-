package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS

// 上述语句用于打包时将指定目录下的文件一并打包

func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		r := gin.Default()
		staticFiles, _ := fs.Sub(FS, "frontend/dist") // 将所有文件打包成一个变量
		// 静态文件都在 /static 这个路由，http.FS 读取文件
		r.StaticFS("/static", http.FS(staticFiles))
		r.NoRoute(func(ctx *gin.Context) {
			path := ctx.Request.URL.Path
			if strings.HasPrefix(path,"/static/"){
				reader,err:=staticFiles.Open("index.html")
				if err!=nil{
					log.Fatal(err)
				}
				defer reader.Close()
				
			}
		})
		r.Run(":8080")
	}()

	// 监听退出信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 先写死路径开启 chrome
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/static/index.html")
	cmd.Start()

	// 等待退出信号
	<-chSignal

	cmd.Process.Kill()
}
