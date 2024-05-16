// Package entity /** 用户表
package entity

type SysUser struct {
	ID           uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"`                  // 主键 用户ID
	UserName     string `json:"user_name" gorm:"not null;unique;type:varchar(255);"`           // 用户名
	Password     string `json:"password" gorm:"type:text;not null;"`                           // 密码 默认值
	Email        string `json:"email" gorm:"not null;unique;"`                                 // 邮箱 不允许为空 唯一
	Phone        string `json:"phone" gorm:"not null;unique;size:11;"`                         // 电话号码 不允许为空 唯一
	Gender       bool   `json:"gender" gorm:"not null;"`                                       // 性别
	LimitType    uint8  `json:"limit_type" gorm:"not null;default:0;"`                         // 用户类型
	Avatar       string `json:"icon" gorm:"default:'http://118.31.43.239:81/image/icon.png';"` // 头像 默认值
	Integral     uint32 `json:"integral" gorm:"not null;default:0;"`                           //用户积分
	Balance      uint32 `json:"balance" gorm:"not null;default:0;"`                            // 余额
	EnableStatus bool   `json:"enable_status" gorm:"not null;default:1;"`                      // 是否激活
	LikeNum      uint32 `json:"like_num" gorm:"not null;default:0;"`                           // 赞
	DontLike     uint32 `json:"dont_like" gorm:"not null;default:0;"`                          // 踩
	UID          string `json:"uid" gorm:"not null;type:varchar(36);"`                         // 用户身份
}

// TableName /** 复写默认方法，设置表名
func (SysUser) TableName() string {
	return "sys_user"
}
