// Package controller
/**
 * @projectName:    Store
 * @package:        controller
 * @className:      VerifyCodeController
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/7 13:28
 * @version:    1.0
 */
package controller

import (
	"Store/request"
	"Store/service"
	"Store/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

// SendEmailCode 获取邮箱验证码
func SendEmailCode(c *gin.Context) {
	// 获取参数接口实例
	var param request.SendEmailParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 调用接口发送验证码
	err, sendStatus := service.SendMsgCodeService(param.Email)
	// 如果发生错误
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 发送失败
	if !sendStatus {
		utils.FailResult(c, "发送失败！")
		return
	}
	// 发送成功
	utils.SuccessResult(c, "发送成功！", nil)
}

// SendImgCode 获取图像验证码
func SendImgCode(c *gin.Context) {
	// 调用接口发送验证码
	err, sendStatus, img := service.SendImgCodeService()
	// 如果发生错误
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 发送失败
	if !sendStatus {
		utils.FailResult(c, "发送失败！")
		return
	}
	imgBase := base64.StdEncoding.EncodeToString(img)
	// 发送成功
	utils.SuccessResult(c, "发送成功！", map[string]string{"img": "data:image/PNG;base64," + imgBase})
}
