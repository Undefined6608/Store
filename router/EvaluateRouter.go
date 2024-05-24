// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      EvaluateRouter
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 20:16
 * @version:    1.0
 */
package routes

import (
	"Store/controller"
	"github.com/gin-gonic/gin"
)

// EvaluateRouter /** 二级路由评论
func EvaluateRouter(router *gin.RouterGroup) {
	// 获取商户评论列表
	router.GET("/getUserEvaluateList", controller.GetUserEvaluateList)
	// 添加商户评论
	router.POST("/addUserEvaluate", controller.AddUserEvaluate)
	// 删除商户评论
	router.POST("/deleteUserEvaluate", controller.DeleteUserEvaluate)
	// 获取商品评论列表
	router.GET("/getProductEvaluateList", controller.GetProductEvaluateList)
	// 添加商品评论
	router.POST("/addProductEvaluate", controller.AddProductEvaluate)
	// 删除商品评论
	router.POST("/deleteProductEvaluate", controller.DeleteProductEvaluate)
}
