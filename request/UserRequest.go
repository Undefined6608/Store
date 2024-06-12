package request

import (
	"github.com/dgrijalva/jwt-go"
)

// UserNameOccupyParams /** 用户名查重参数
type UserNameOccupyParams struct {
	UserName string `json:"userName" binding:"required,min=3,max=15"`
}

// PhoneOccupyParams /** 电话号码查重参数
type PhoneOccupyParams struct {
	Phone string `json:"phone" binding:"required,len=11"`
}

// EmailOccupyParams /** 邮箱查重参数
type EmailOccupyParams struct {
	Email string `json:"email" binding:"required"`
}

// SendEmailParams /** 邮箱查重参数
type SendEmailParams struct {
	Email string `json:"email" binding:"required"`
}

// RegisterParams /** 注册参数
type RegisterParams struct {
	UserName    string `json:"userName" binding:"required,min=3,max=15"`
	Phone       string `json:"phone" binding:"required,len=11"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar" binding:"required"`
	EmailCode   string `json:"emailCode" binding:"required,len=6"`
	ImgCode     string `json:"imgCode" binding:"required,len=6"`
	Gender      int8   `json:"gender" binding:"required"`
	LimitType   uint8  `json:"limitType" binding:"required"`
	Password    string `json:"password" binding:"required,len=32"`
	VerPassword string `json:"verPassword" binding:"required,len=32"`
}

// PhoneLoginParams /** 电话号码登录参数
type PhoneLoginParams struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,len=32"`
	ImgCode  string `json:"imgCode" binding:"required,len=6"`
}

// TokenInfo /** Token信息
type TokenInfo struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"`                  // 主键 用户ID
	UserName  string `json:"userName" gorm:"not null;unique;type:varchar(255);"`            // 用户名
	Email     string `json:"email" gorm:"not null;unique;"`                                 // 邮箱 不允许为空 唯一
	Phone     string `json:"phone" gorm:"not null;unique;size:11;"`                         // 电话号码 不允许为空 唯一
	Gender    bool   `json:"gender" gorm:"not null;"`                                       // 性别
	LimitType uint8  `json:"limitType" gorm:"not null;default:0;"`                          // 用户类型
	Avatar    string `json:"icon" gorm:"default:'http://118.31.43.239:81/image/icon.png';"` // 头像 默认值
	Integral  uint32 `json:"integral" gorm:"not null;default:0;"`                           //用户积分
	Balance   uint32 `json:"balance" gorm:"not null;default:0;"`                            // 余额
	LikeNum   uint32 `json:"likeNum" gorm:"not null;default:0;"`                            // 赞
	DontLike  uint32 `json:"dontLike" gorm:"not null;default:0;"`                           // 踩
	UID       string `json:"uid" gorm:"not null;type:varchar(36);"`                         // 用户身份
}

// TokenParams /** 定义 Token 类型
type TokenParams struct {
	UserInfo           TokenInfo // 用户信息
	jwt.StandardClaims           // token 配置
}

// EmailLoginParams /** 邮箱登录参数
type EmailLoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,len=32"`
	ImgCode  string `json:"imgCode" binding:"required,len=6"`
}

// ForgotPasswordParams /** 忘记密码参数
type ForgotPasswordParams struct {
	Email       string `json:"email" binding:"required"`
	EmailCode   string `json:"emailCode" binding:"required,len=6"`
	ImgCode     string `json:"imgCode" binding:"required,len=6"`
	NewPassword string `json:"newPassword" binding:"required,len=32"`
	VerPassword string `json:"verPassword" binding:"required,len=32"`
}

// ModifyUserInfoParams /** 修改用户信息参数
type ModifyUserInfoParams struct {
	UserName  string `json:"userName" binding:"required,min=3,max=15"`
	Phone     string `json:"phone" binding:"required,len=11"`
	Email     string `json:"email" binding:"required"`
	Gender    uint8  `json:"gender" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
	ImgCode   string `json:"imgCode" binding:"required,len=6"`
	EmailCode string `json:"emailCode" binding:"required,len=6"`
}

// ModifyPasswordParams /** 修改密码参数
type ModifyPasswordParams struct {
	OldPassword      string `json:"oldPassword" binding:"required,len=32"`
	NewPassword      string `json:"newPassword" binding:"required,len=32"`
	ImgCode          string `json:"imgCode" binding:"required,len=6"`
	VerifiedPassword string `json:"verifiedPassword" binding:"required,len=32"`
}

// DeleteUserParams /** 删除用户
type DeleteUserParams struct {
	UserUid string `json:"userUid" binding:"required"`
}

// GetUserNameByUidParam /** 通过 Uid 获取用户名
type GetUserNameByUidParam struct {
	UserUid string `form:"userUid" binding:"required"`
}
