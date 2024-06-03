// Package controller
/**
 * @projectName:    Store
 * @package:        controller
 * @className:      AddressController
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 17:08
 * @version:    1.0
 */
package controller

import (
	"Store/entity"
	"Store/request"
	"Store/service"
	"Store/utils"
	"github.com/gin-gonic/gin"
)

// AddUserAddress /** 添加用户地址
func AddUserAddress(c *gin.Context) {
	var param request.AddUserAddressParams
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, userInfo := utils.GetCacheUser(c)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if err := service.AddUserAddressService(&param, userInfo); err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "新增成功", nil)
}

// GetUserAddressList /** 获取用户地址列表
func GetUserAddressList(c *gin.Context) {
	err, params := utils.GetCacheUser(c)
	err, addresses := service.GetUserAddressListService(params)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.UserAddress{"userAddress": addresses})
}

// ModifyUserAddress /** 修改用户地址
func ModifyUserAddress(c *gin.Context) {
	var param request.ModifyUserAddressParams
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.ModifyUserAddressService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "修改失败")
		return
	}
	utils.SuccessResult(c, "修改成功", nil)
}

// DeleteUserAddress /** 删除用户地址
func DeleteUserAddress(c *gin.Context) {
	var param request.DeleteUserAddressParams
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.DeleteUserAddressService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "删除失败")
		return
	}
	utils.SuccessResult(c, "删除成功", nil)
}

// AddOrderForm /** 添加订单
func AddOrderForm(c *gin.Context) {
	var param request.AddOrderFormParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.AddOrderFormService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "新增失败")
		return
	}
	utils.SuccessResult(c, "新增成功", nil)
}

// GetUserOrderFormList /** 获取订单列表
func GetUserOrderFormList(c *gin.Context) {
	err, params := utils.GetCacheUser(c)
	err, orderFormList := service.GetUserOrderFormListService(params)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.OrderForm{"orderFormList": orderFormList})
}

// GetOrderFormInfo /** 获取订单详情
func GetOrderFormInfo(c *gin.Context) {
	var param request.GetOrderFormInfoParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, orderForm := service.GetOrderFormInfoService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string]entity.OrderForm{"orderForm": orderForm})
}

// ModifyOrderForm /** 修改订单
func ModifyOrderForm(c *gin.Context) {
	var param request.ModifyOrderFormParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.ModifyOrderFormService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "修改失败")
		return
	}
	utils.SuccessResult(c, "修改成功", nil)
}

// DeleteOrderForm /** 删除订单
func DeleteOrderForm(c *gin.Context) {
	var param request.DeleteOrderFormParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.DeleteOrderFormService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "删除失败")
		return
	}
	utils.SuccessResult(c, "删除成功", nil)
}

// SetOrderFormStatus /** 修改订单状态
func SetOrderFormStatus(c *gin.Context) {
	var param request.SetOrderFormStatusParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.SetOrderFormStatusService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "修改失败")
		return
	}
	utils.SuccessResult(c, "修改成功", nil)
}
