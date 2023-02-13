package main

import (
	"MinimalistTiktok/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	_, err := os.Stat(config.VideosImagePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(config.VideosImagePath, os.ModePerm)
			if err != nil {
				fmt.Println("创建文件夹失败", err)
				return
			}
		}
	}
}

func main() {
	Init()
	r := gin.Default()
	// 使用权限验证中间件
	//r.Use(middleware.Auth())
	InitRouter(r)
	err := r.Run(":" + config.ServerPort)
	if err != nil {
		return
	}
}
