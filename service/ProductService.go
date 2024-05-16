// Package service
/**
 * @projectName:    Store
 * @package:        service
 * @className:      ProductService
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 12:50
 * @version:    1.0
 */
package service

import (
	"Store/entity"
	"Store/request"
	"Store/utils"
	"errors"
)

// GetBannerListService 获取轮播图列表
func GetBannerListService() (error, []entity.ProductBanner) {
	var banner []entity.ProductBanner
	pool.Model(&entity.ProductBanner{}).Find(&banner)
	if len(banner) == 0 {
		return errors.New("暂无轮播图"), nil
	}
	return nil, banner
}

// GetProductListService 获取商品列表
func GetProductListService(param *request.GetProductByTypeParam) (error, []entity.Product) {
	var productList []entity.Product
	if param.TypeId == 0 {
		pool.Find(&productList)
		if len(productList) == 0 {
			return errors.New("暂无商品列表"), nil
		}
		return nil, productList
	}
	pool.Model(&entity.Product{}).Where("type_id=?", param.TypeId).Find(&productList)
	if len(productList) == 0 {
		return errors.New("暂无商品列表"), nil
	}
	return nil, productList
}

// GetProductInfoService 获取商品详情
func GetProductInfoService(param *request.GetProductInfoParam) (error, entity.Product) {
	var productInfo entity.Product
	pool.Model(&entity.Product{}).Where("id=?", param.ProductId).First(&productInfo)
	if productInfo.Price == 0 {
		return errors.New("暂未查到"), productInfo
	}
	return nil, productInfo
}

func AddProductTypeService(param *request.AddProductTypeParam) (error, bool) {

	return nil, true
}

func ModifyProductTypeService(param *request.ModifyProductTypeParam) (error, bool) {
	return nil, true
}

func GetProductTypeService() (error, []entity.ProductType) {
	var productType []entity.ProductType
	pool.Model(&entity.ProductType{}).Find(&productType)
	if len(productType) == 0 {
		return errors.New("暂无数据"), nil
	}
	return nil, productType
}

// AddProductService 添加商品
func AddProductService(tokenInfo *request.TokenParams, param *request.AddProductParam) (error, bool) {
	var productType entity.ProductType
	var product entity.Product
	var verProduct entity.Product
	var productBanner entity.ProductBanner
	// 判断类型是否存在
	if err := pool.Model(&entity.ProductType{}).Where("id=?", param.TypeId).First(&productType).Error; err != nil {
		return err, false
	}
	// 设置product的值
	product.UID = utils.CreateUUID()
	product.Name = param.Name
	product.TypeId = param.TypeId
	product.Icon = param.Icon
	product.Price = param.Price
	product.Content = param.Content
	product.UserUId = tokenInfo.UserInfo.UID

	if err := pool.Create(&product).Error; err != nil {
		return err, false
	}
	if err := pool.Where("uid=?", product.UID).First(&verProduct).Error; err != nil {
		return err, false
	}
	productBanner.ProductID = verProduct.ID
	productBanner.Banner = param.Banner
	productBanner.ProductName = param.Name
	productBanner.UID = utils.CreateUUID()
	if err := pool.Create(&productBanner).Error; err != nil {
		return err, false
	}
	return nil, true
}

// ModifyProductService 修改商品
func ModifyProductService(param *request.ModifyProductParam) (error, bool) {
	return nil, true
}

// DeleteProductService 删除商品
func DeleteProductService(param *request.DeleteProductParam) (error, bool) {
	return nil, true
}

// SetProductStatusService 设置商品是否上架
func SetProductStatusService(param *request.SetProductStatusParam) (error, bool) {
	return nil, true
}

func DeleteProductTypeService(param *request.DeleteProductTypeParam) (error, bool) {
	return nil, true
}
