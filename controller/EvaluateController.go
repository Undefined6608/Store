// Package controller
/**
 * @projectName:    Store
 * @package:        controller
 * @className:      EvaluateController
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 20:03
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

// AddUserEvaluate /** 添加商户评论
func AddUserEvaluate(c *gin.Context) {
	var param request.AddUserEvaluateRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.AddUserEvaluateService(&param)
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

// GetUserEvaluateList /** 获取商户评论列表
func GetUserEvaluateList(c *gin.Context) {
	var param request.GetUserEvaluateRequest
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, userEvaluateList := service.GetUserEvaluateListService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.UserEvaluate{"userEvaluateList": userEvaluateList})
}

// DeleteUserEvaluate /** 删除商户评论
func DeleteUserEvaluate(c *gin.Context) {
	var param request.DeleteUserEvaluateRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.DeleteUserEvaluateService(&param)
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

// AddProductEvaluate /** 添加商品评论
func AddProductEvaluate(c *gin.Context) {
	var param request.AddProductEvaluateRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.AddProductEvaluateService(&param)
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

// GetProductEvaluateList /** 获取商品评论列表
func GetProductEvaluateList(c *gin.Context) {
	var param request.GetProductEvaluateRequest
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, userEvaluateList := service.GetProductEvaluateListService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.ProductEvaluate{"productEvaluateList": userEvaluateList})
}

// DeleteProductEvaluate /** 删除商品评论
func DeleteProductEvaluate(c *gin.Context) {
	var param request.DeleteProductEvaluateRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.DeleteProductEvaluateService(&param)
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
