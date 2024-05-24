// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      AddressRouter
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 17:11
 * @version:    1.0
 */
package routes

import (
	"Store/controller"
	"github.com/gin-gonic/gin"
)

// OrderFormRouter /** 二级路由订单
func OrderFormRouter(router *gin.RouterGroup) {
	// 获取用户地址列表
	router.GET("/getUserAddressList", controller.GetUserAddressList)
	// 添加用户地址
	router.POST("/addUserAddress", controller.AddUserAddress)
	// 修改用户地址
	router.POST("/modifyUserAddress", controller.ModifyUserAddress)
	// 删除用户地址
	router.POST("/deleteUserAddress", controller.DeleteUserAddress)
	// 获取订单列表
	router.GET("/getOrderFormList", controller.GetUserOrderFormList)
	// 获取订单详情
	router.GET("/getOrderFormInfo", controller.GetOrderFormInfo)
	// 添加订单
	router.POST("/addOrderForm", controller.AddOrderForm)
	// 修改订单
	router.POST("/modifyOrderForm", controller.ModifyOrderForm)
	// 删除订单
	router.POST("/deleteOrderForm", controller.DeleteOrderForm)
	// 修改订单状态
	router.POST("/setOrderFormStatus", controller.SetOrderFormStatus)
}
