// Package service
/**
 * @projectName:    Store
 * @package:        service
 * @className:      EvaluateService
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 19:54
 * @version:    1.0
 */
package service

import (
	"Store/entity"
	"Store/request"
)

// GetUserEvaluateListService /** 获取商户评论列表
func GetUserEvaluateListService(param *request.GetUserEvaluateRequest) (error, []entity.UserEvaluate) {
	var userEvaluateList []entity.UserEvaluate
	pool.Model(&entity.UserEvaluate{}).Where("merchantId=?", param.MerchantId).Find(&userEvaluateList)
	return nil, userEvaluateList
}

// AddUserEvaluateService /** 添加商户评论
func AddUserEvaluateService(param *request.AddUserEvaluateRequest) (error, bool) {
	return nil, true
}

// DeleteUserEvaluateService /** 删除商户评论
func DeleteUserEvaluateService(param *request.DeleteUserEvaluateRequest) (error, bool) {
	return nil, true
}

// GetProductEvaluateListService /** 获取商品评论列表
func GetProductEvaluateListService(param *request.GetProductEvaluateRequest) (error, []entity.ProductEvaluate) {
	var productEvaluateList []entity.ProductEvaluate
	pool.Model(&entity.ProductEvaluate{}).Where("product_id=?", param.ProductId).Find(&productEvaluateList)
	return nil, productEvaluateList
}

// AddProductEvaluateService /** 添加商品评论
func AddProductEvaluateService(param *request.AddProductEvaluateRequest) (error, bool) {
	return nil, true
}

// DeleteProductEvaluateService /** 删除商品评论
func DeleteProductEvaluateService(param *request.DeleteProductEvaluateRequest) (error, bool) {
	return nil, true
}
