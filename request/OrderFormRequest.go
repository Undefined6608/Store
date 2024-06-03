// Package request
/**
 * @projectName:    Store
 * @package:        request
 * @className:      AddressRequest
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 17:09
 * @version:    1.0
 */
package request

// AddUserAddressParams /** 添加用户地址
type AddUserAddressParams struct {
	Consignee string `json:"consignee" binding:"required"` //收货人
	Phone     string `json:"phone" binding:"required"`     // 收货人联系电话
	Gender    int8   `json:"gender" binding:"required"`    // 性别
	Address   string `json:"address" binding:"required"`   // 收货地址
	Def       int8   `json:"def" binding:"required"`       // 是否为默认
}

// ModifyUserAddressParams /** 修改用户地址
type ModifyUserAddressParams struct {
	Id        uint32 `json:"id" binding:"required"`        // 用户地址ID
	Consignee string `json:"consignee" binding:"required"` //收货人
	Phone     string `json:"phone" binding:"required"`     // 收货人联系电话
	Gender    bool   `json:"gender" binding:"required"`    // 性别
	Address   string `json:"address" binding:"required"`   // 收货地址
	Def       bool   `json:"def" binding:"required"`       // 是否为默认
}

// DeleteUserAddressParams /** 删除用户地址
type DeleteUserAddressParams struct {
	AddressId uint32 `json:"addressId" binding:"required"` // 用户地址ID
}

// AddOrderFormParam /** 添加订单
type AddOrderFormParam struct {
	ProductId uint32 `json:"product_id" binding:"required"` // 商品ID
	ToAddress uint32 `json:"to_address" binding:"required"` // 送货地址ID
	Remarks   string `json:"remarks" binding:"required"`    // 订单备注
}

// GetOrderFormInfoParam /** 获取订单信息
type GetOrderFormInfoParam struct {
	OderFormId uint32 `json:"oderFormId" binding:"required"` // 订单ID
}

// ModifyOrderFormParam /** 修改订单
type ModifyOrderFormParam struct {
	OderFormId uint32 `json:"oderFormId" binding:"required"` // 订单ID
	ToAddress  uint32 `json:"to_address" binding:"required"` // 送货地址ID
	Remarks    string `json:"remarks" binding:"required"`    // 订单备注
}

// DeleteOrderFormParam /** 删除订单
type DeleteOrderFormParam struct {
	OderFormId uint32 `json:"oderFormId" binding:"required"` // 订单ID
}

// SetOrderFormStatusParam /** 设置订单状态
type SetOrderFormStatusParam struct {
	OderFormId uint32 `json:"oderFormId" binding:"required"` // 订单ID
	Status     uint8  `json:"status" binding:"required"`     //订单状态
}
