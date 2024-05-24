// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      VerifyCodeRouter
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/7 13:35
 * @version:    1.0
 */
package routes

import (
	"Store/controller"
	"github.com/gin-gonic/gin"
)

// VerifyCodeRouter /** 验证码路由
func VerifyCodeRouter(router *gin.RouterGroup) {
	// 邮箱验证码
	router.POST("/emailCode", controller.SendEmailCode)
	// 图像验证码
	router.GET("/imgCode", controller.SendImgCode)
}
