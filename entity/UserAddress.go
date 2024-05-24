// Package entity /** 用户地址
package entity

type UserAddress struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 主键 地址ID
	UserUid   string `json:"user_uid" gorm:"not null;"`                    // 用户UID
	Consignee string `json:"consignee" gorm:"not null;"`                   //收货人
	Phone     string `json:"phone" gorm:"not null;"`                       // 收货人联系电话
	Gender    bool   `json:"gender" gorm:"not null;default:0;"`            // 性别
	Address   string `json:"address" gorm:"not null;"`                     // 收货地址
	Def       bool   `json:"def" gorm:"not null;default:0"`                // 是否为默认
}

// TableName /** 复写默认方法，设置表名
func (UserAddress) TableName() string {
	return "user_address"
}
