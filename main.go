package main

import (
	"TVHelper/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routers.LoadDouBan(r)
	routers.LoadConfig(r)
	routers.LoadLive(r)

	if err := r.Run(":16214"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
