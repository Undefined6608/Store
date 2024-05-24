// Package routes
/**
 * @projectName:    Store
 * @package:        routes
 * @className:      UserRouter
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/7 12:34
 * @version:    1.0
 */
package routes

import (
	"Store/controller"
	"github.com/gin-gonic/gin"
)

// UserRouter /** 二级路由用户
func UserRouter(router *gin.RouterGroup) {
	// 用户测试接口
	router.POST("/helloUser", controller.HelloUser)
	// 用户名查重
	router.POST("/userNameOccupy", controller.UserNameOccupy)
	// 电话号码查重
	router.POST("/phoneOccupy", controller.PhoneOccupy)
	// 邮箱查重
	router.POST("/emailOccupy", controller.EmailOccupy)
	// 注册
	router.POST("/register", controller.Register)
	// 电话号码登录
	router.POST("/phoneLogin", controller.PhoneLogin)
	// 邮箱登录
	router.POST("/emailLogin", controller.EmailLogin)
	// 忘记密码
	router.POST("/forgotPassword", controller.ForgotPassword)
	// 查询用户信息
	router.GET("/userInfo", controller.UserInfo)
	// 获取全部用户信息
	router.GET("/allUserInfo", controller.AllUserInfo)
	// 修改用户信息
	router.POST("/modifyUserInfo", controller.ModifyUserInfo)
	// 修改密码
	router.POST("/modifyPassword", controller.ModifyPassword)
	// 退出登录
	router.POST("/logout", controller.Logout)
}
