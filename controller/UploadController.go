package controller

import (
	"Store/config"
	"Store/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// UserAvatar /** 上传用户头像
func UserAvatar(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}

	// 生成 UUID 作为文件名的一部分
	newName := utils.GenerateFileName(file.Filename)

	// 检查文件类型是否被允许上传
	ext := filepath.Ext(file.Filename)
	if !utils.IsAllowedImageType(ext) {
		utils.FailResult(c, "只允许上传图片文件")
		return
	}

	// 检查文件大小
	if file.Size > (config.Upload.MaxSize.Img << 20) { // 限制文件大小
		utils.FailResult(c, "文件大小超过限制")
		return
	}

	// 保存上传的文件到本地
	dst := fmt.Sprintf(config.Upload.ImgLoad.User+"%s", newName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		utils.FailResult(c, err.Error())
		return
	}

	utils.SuccessResult(c, "上传成功", map[string]interface{}{"url": config.Upload.Host + "icon/" + newName})
}

// ProductIcon /** 上传商品图像
func ProductIcon(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}

	// 生成 UUID 作为文件名的一部分
	newName := utils.GenerateFileName(file.Filename)

	// 检查文件类型是否被允许上传
	ext := filepath.Ext(file.Filename)
	if !utils.IsAllowedImageType(ext) {
		utils.FailResult(c, "只允许上传图片文件")
		return
	}

	// 检查文件大小
	if file.Size > (config.Upload.MaxSize.Img << 20) { // 限制文件大小
		utils.FailResult(c, "文件大小超过限制")
		return
	}

	// 保存上传的文件到本地
	dst := fmt.Sprintf(config.Upload.ImgLoad.Product+"%s", newName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		utils.FailResult(c, err.Error())
		return
	}

	utils.SuccessResult(c, "上传成功", map[string]interface{}{"url": config.Upload.Host + "product/" + newName})
}
