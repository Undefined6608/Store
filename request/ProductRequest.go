// Package request
/**
 * @projectName:    Store
 * @package:        request
 * @className:      ProductRequest
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 14:36
 * @version:    1.0
 */
package request

// GetProductTypeParam /** 获取商品类型
type GetProductTypeParam struct {
	TypeId uint32 `form:"typeId"`
}

// GetProductByTypeParam /** 通过类型查询商品
type GetProductByTypeParam struct {
	TypeId  uint32 `form:"typeId"`  // 类型ID
	UserUid string `form:"userUid"` // 用户ID
}

// GetProductInfoParam /** 获取商品信息
type GetProductInfoParam struct {
	ProductId uint32 `form:"productId" binding:"required"` // 商品ID
}

// AddProductParam /** 添加商品
type AddProductParam struct {
	TypeId  uint32 `json:"typeId" binding:"required"`  // 商品类型ID
	Name    string `json:"name" binding:"required"`    // 商品名字
	Banner  string `json:"banner" binding:"required"`  // 商品轮播图
	Icon    string `json:"icon" binding:"required"`    // 商品图标
	Price   uint32 `json:"price" binding:"required"`   // 商品价格
	Content string `json:"content" binding:"required"` // 商品介绍
}

// ModifyProductParam /** 修改商品
type ModifyProductParam struct {
	Id      uint32 `json:"id" binding:"required"`      // 商品ID
	TypeId  uint32 `json:"typeId" binding:"required"`  // 商品类型ID
	Banner  string `json:"banner" binding:"required"`  // 商品轮播图
	Icon    string `json:"icon" binding:"required"`    // 商品图标
	Price   uint32 `json:"price" binding:"required"`   // 商品价格
	Content string `json:"content" binding:"required"` // 商品介绍
}

// DeleteProductParam /** 删除商品
type DeleteProductParam struct {
	ProductID uint32 `json:"productId" binding:"required"` // 商品ID
}

// SetProductStatusParam /** 设置商品状态
type SetProductStatusParam struct {
	ProductID uint32 `json:"productId" binding:"required"` // 商品ID
}

// AddProductTypeParam /** 添加商品类型
type AddProductTypeParam struct {
}

// ModifyProductTypeParam /** 修改商品类型
type ModifyProductTypeParam struct {
}

// DeleteProductTypeParam /** 删除商品类型
type DeleteProductTypeParam struct {
}
