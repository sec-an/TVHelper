package main

import (
	"TVHelper/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routers.LoadDouBan(r) // 豆瓣主页
	routers.LoadConfig(r) // 代理配置
	routers.LoadLive(r)   // 直播文件

	if err := r.Run(":16214"); err != nil {
		log.Fatalf("startup service failed: %v\n\n", err)
	}
}
