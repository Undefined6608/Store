// Package service
/**
 * @projectName:    Store
 * @package:        service
 * @className:      FeedBackService
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/5/24 9:08
 * @version:    1.0
 */
package service

import (
	"Store/entity"
	"Store/request"
	"errors"
)

// AddFeedBackService /** 添加反馈信息
func AddFeedBackService(param *request.AddFeedbackRequest) error {
	if err := pool.Create(&entity.FeedBack{
		Name:    param.Name,
		Phone:   param.Phone,
		Email:   param.Email,
		Content: param.Content,
	}).Error; err != nil {
		return errors.New("添加失败！")
	}
	return nil
}

// GetFeedBackService /** 获取反馈信息
func GetFeedBackService(param *request.GetFeedbackRequest, userInfo *request.TokenParams) (error, []entity.FeedBack) {
	var feedBackList []entity.FeedBack
	if userInfo.UserInfo.LimitType != 1 {
		return errors.New("权限错误！"), nil
	}
	err := pool.Find(&feedBackList).Error
	if err != nil {
		return errors.New("获取失败！"), nil
	}
	return nil, feedBackList
}
