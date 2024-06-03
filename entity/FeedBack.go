// Package entity
/**
 * @projectName:    Store
 * @package:        entity
 * @className:      FeedBack
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/5/24 9:03
 * @version:    1.0
 */
package entity

type FeedBack struct {
	ID      uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 主键 商品路径ID
	Name    string `json:"name" gorm:"not null;type:varchar(255)"`      // 姓名
	Phone   string `json:"phone" gorm:"not null;type:varchar(255)"`     // 电话号码
	Email   string `json:"email" gorm:"not null;type:varchar(255)"`     // 邮箱
	Content string `json:"content" gorm:"not null;type:varchar(255)"`   // 反馈内容
}

// TableName /** 复写默认方法，设置表名
func (FeedBack) TableName() string {
	return "feed_back"
}
