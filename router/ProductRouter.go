// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      ProductRouter
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 14:31
 * @version:    1.0
 */
package routes

import (
	"Store/controller"
	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.RouterGroup) {
	// 获取轮播图
	router.GET("/getBannerList", controller.GetBannerList)
	// 获取商品列表
	router.GET("/getProductList", controller.GetProductList)
	// 获取商品详情
	router.GET("/getProductInfo", controller.GetProductInfo)
	// 添加商品
	router.POST("/addProduct", controller.AddProduct)
	// 修改商品
	router.POST("/modifyProduct", controller.ModifyProduct)
	// 删除商品
	router.POST("/deleteProduct", controller.DeleteProduct)
	// 设置商品状态
	router.POST("/setProductStatus", controller.SetProductStatus)
	// 查询商品类型
	router.GET("/getProductType", controller.GetProductType)
	// 添加商品类型
	router.POST("/addProductType", controller.AddProductType)
	// 修改商品类型
	router.POST("/modifyProductType", controller.ModifyProductType)
	// 删除商品类型
	router.POST("/deleteProductType", controller.DeleteProductType)
}
