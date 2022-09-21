package controller

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func AddressesController(c *gin.Context) {
	add, _ := net.InterfaceAddrs()
	var result []string
	for _, address := range add {
		// 类型断言
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": result,
	})
}
