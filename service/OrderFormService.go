// Package service
/**
 * @projectName:    Store
 * @package:        service
 * @className:      AddressService
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 17:10
 * @version:    1.0
 */
package service

import (
	"Store/entity"
	"Store/request"
	"errors"
)

// GetUserAddressListService /** 获取用户地址列表
func GetUserAddressListService(param *request.TokenParams) (error, []entity.UserAddress) {
	var addressList []entity.UserAddress
	pool.Model(&entity.UserAddress{}).Where("user_uid=?", param.UserInfo.UID).Find(&addressList)
	return nil, addressList
}

// AddUserAddressService /** 添加用户地址
func AddUserAddressService(param *request.AddUserAddressParams, userInfo *request.TokenParams) error {
	err := pool.Model(&entity.UserAddress{}).Create(&entity.UserAddress{
		UserUid:   userInfo.UserInfo.UID,
		Consignee: param.Consignee,
		Phone:     param.Phone,
		Gender:    param.Gender == 2,
		Address:   param.Address,
		Def:       param.Def == 2,
	}).Error
	if err != nil {
		return errors.New("新增失败！")
	}
	return nil
}

// ModifyUserAddressService /** 修改用户地址
func ModifyUserAddressService(param *request.ModifyUserAddressParams) (error, bool) {
	return nil, true
}

// DeleteUserAddressService /** 删除商品
func DeleteUserAddressService(param *request.DeleteUserAddressParams) (error, bool) {
	return nil, true
}

// GetUserOrderFormListService /** 获取订单列表
func GetUserOrderFormListService(param *request.TokenParams) (error, []entity.OrderForm) {
	var orderForm []entity.OrderForm
	pool.Model(&entity.OrderForm{}).Where("user_uid=?", param.UserInfo.UID).Find(&orderForm)
	return nil, orderForm
}

// GetOrderFormInfoService /** 获取订单详情
func GetOrderFormInfoService(param *request.GetOrderFormInfoParam) (error, entity.OrderForm) {
	var orderForm entity.OrderForm
	pool.Model(&entity.OrderForm{}).Where("id=?", param.OderFormId).Find(&orderForm)
	return nil, orderForm
}

// AddOrderFormService /** 添加订单
func AddOrderFormService(param *request.AddOrderFormParam) (error, bool) {
	return nil, true
}

// ModifyOrderFormService /** 修改订单
func ModifyOrderFormService(param *request.ModifyOrderFormParam) (error, bool) {
	return nil, true
}

// DeleteOrderFormService /** 删除订单
func DeleteOrderFormService(param *request.DeleteOrderFormParam) (error, bool) {
	return nil, true
}

// SetOrderFormStatusService /** 修改订单状态
func SetOrderFormStatusService(param *request.SetOrderFormStatusParam) (error, bool) {
	return nil, true
}
