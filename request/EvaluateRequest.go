// Package request
/**
 * @projectName:    Store
 * @package:        request
 * @className:      EvaluateRequest
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 19:46
 * @version:    1.0
 */
package request

// GetUserEvaluateRequest /** 获取用户评论信息
type GetUserEvaluateRequest struct {
	MerchantId uint32 `form:"merchantId" binding:"required"`
}

// AddUserEvaluateRequest /** 添加用户评论
type AddUserEvaluateRequest struct {
}

// DeleteUserEvaluateRequest /** 删除用户评论
type DeleteUserEvaluateRequest struct {
}

// GetProductEvaluateRequest /** 获取商品评论
type GetProductEvaluateRequest struct {
	ProductId uint32 `form:"productId" binding:"required"`
}

// AddProductEvaluateRequest /** 添加商品评论
type AddProductEvaluateRequest struct {
}

// DeleteProductEvaluateRequest /** 删除商品评论
type DeleteProductEvaluateRequest struct {
}
