package main

import (
	"Store/config"
	"Store/middleware"
	routes "Store/router"
	"Store/service"
	"fmt"
	gonic "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 获取项目路由
	router := gonic.New()
	// 挂载中间件
	router.Use(gonic.Logger(), gonic.Recovery(), middleware.CorsMiddleware(), middleware.LoggerToFile(), middleware.JwtVerifyMiddle())
	// 加载代理中间件
	err := router.SetTrustedProxies([]string{"192.168.1.0/24"})
	if err != nil {
		fmt.Println("代理失败！")
		return
	}

	// 验证数据表是否存在
	service.VerDataBase()
	// 编写项目基础接口
	router.GET("/test", func(c *gonic.Context) {
		c.JSON(http.StatusOK, "HelloWorld!")
	})
	// 调用项目主路由
	routes.SetupRouterGroup(router.Group("/api"))
	// 开启端口监听
	err = router.Run(":" + config.Default().Port)
	// 开启监听失败
	if err != nil {
		// 写入日志
		log.Fatalln("项目启动失败!")
		return
	}
}
