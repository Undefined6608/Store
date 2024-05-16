// Package controller
/**
 * @projectName:    Store
 * @package:        controller
 * @className:      ProductController
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 12:37
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

// GetBannerList /** 获取轮播列表
func GetBannerList(c *gin.Context) {
	err, banners := service.GetBannerListService()
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.ProductBanner{"bannerList": banners})
}

// GetProductList /** 获取商品列表
func GetProductList(c *gin.Context) {
	var param request.GetProductByTypeParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, products := service.GetProductListService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.Product{"productList": products})
}

// GetProductInfo /** 获取商品详细信息
func GetProductInfo(c *gin.Context) {
	var param request.GetProductInfoParam

	err := c.ShouldBindQuery(&param)

	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}

	err, product := service.GetProductInfoService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string]entity.Product{"productInfo": product})
}

// AddProduct /** 添加商品
func AddProduct(c *gin.Context) {
	var param request.AddProductParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, tokenInfo := utils.GetCacheUser(c)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.AddProductService(tokenInfo, &param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "添加失败")
		return
	}
	utils.SuccessResult(c, "添加成功！", nil)
}

// ModifyProduct /** 修改商品信息
func ModifyProduct(c *gin.Context) {
	var param request.ModifyProductParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.ModifyProductService(&param)
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

// DeleteProduct /** 删除商品
func DeleteProduct(c *gin.Context) {
	var param request.DeleteProductParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.DeleteProductService(&param)
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

// SetProductStatus /** 设置商品状态
func SetProductStatus(c *gin.Context) {
	var param request.SetProductStatusParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.SetProductStatusService(&param)
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

// GetProductType /** 获取商品类型
func GetProductType(c *gin.Context) {
	err, products := service.GetProductTypeService()
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.ProductType{"productType": products})
}

// AddProductType /** 添加商品类型
func AddProductType(c *gin.Context) {
	var param request.AddProductTypeParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}

	err, status := service.AddProductTypeService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "添加失败")
		return
	}
	utils.SuccessResult(c, "添加成功！", nil)
}

// ModifyProductType /** 修改商品类型
func ModifyProductType(c *gin.Context) {
	var param request.ModifyProductTypeParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.ModifyProductTypeService(&param)
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

// DeleteProductType /** 删除商品类型
func DeleteProductType(c *gin.Context) {
	var param request.DeleteProductTypeParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, status := service.DeleteProductTypeService(&param)
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
