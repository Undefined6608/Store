// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      index
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/7 12:33
 * @version:    1.0
 */
package routes

import "github.com/gin-gonic/gin"

// SetupRouterGroup /** 项目主路由（一级）
func SetupRouterGroup(router *gin.RouterGroup) {
	// 调用用户路由
	UserRouter(router.Group("/user"))
	// 调用上传路由
	UploadRouter(router.Group("/upload"))
	// 验证码路由
	VerifyCodeRouter(router.Group("/verifyCode"))
	// 商品路由
	ProductRouter(router.Group("/product"))
	// 地址路由
	OrderFormRouter(router.Group("/orderForm"))
	// 评论路由
	EvaluateRouter(router.Group("/evaluate"))
	// 反馈路由
	FeedBackRouter(router.Group("/feedBack"))
}
