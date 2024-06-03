// Package controller
/**
 * @projectName:    Store
 * @package:        controller
 * @className:      FeedBackController
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/5/24 9:07
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

// AddFeedBack /** 添加反馈
func AddFeedBack(c *gin.Context) {
	var param request.AddFeedbackRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err = service.AddFeedBackService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "添加成功！", nil)
}

// GetFeedBack /** 获取反馈
func GetFeedBack(c *gin.Context) {
	var param request.GetFeedbackRequest
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	err, userInfo := utils.GetCacheUser(c)
	err, list := service.GetFeedBackService(&param, userInfo)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	utils.SuccessResult(c, "获取成功！", map[string][]entity.FeedBack{"feedbackList": list})
}
