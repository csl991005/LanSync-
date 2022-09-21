package server

import (
	"LanSync/server/controller"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS

// 上述语句用于打包时将指定目录下的文件一并打包

func Run() {
	port := "27149"
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	staticFiles, _ := fs.Sub(FS, "frontend/dist") // 将所有文件打包成一个变量
	// 静态文件都在 /static 这个路由，http.FS 读取文件
	r.StaticFS("/static", http.FS(staticFiles))
	r.POST("/api/v1/files", controller.FilesController)
	r.GET("/api/v1/qrcodes", controller.QrcodesController)
	r.GET("/uploads/:path", controller.UploadsController)
	r.GET("/api/v1/addresses", controller.AddressesController)
	r.POST("/api/v1/texts", controller.TextController)
	r.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			ctx.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	})
	r.Run(":" + port)
}
