// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      FeedBackRouter
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/5/24 9:11
 * @version:    1.0
 */
package routes

import (
	"Store/controller"
	"github.com/gin-gonic/gin"
)

func FeedBackRouter(router *gin.RouterGroup) {
	// 添加反馈
	router.POST("/addFeedBack", controller.AddFeedBack)
	// 获取反馈信息
	router.GET("/getFeedBack", controller.GetFeedBack)
}
