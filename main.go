package main

import (
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 使用权限验证中间件
	r.Use(middleware.Auth())
	initRouter(r)
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
