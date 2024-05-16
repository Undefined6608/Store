// Package entity /** 用户评论
package entity

type UserEvaluate struct {
	ID            uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 主键 评论ID
	MerchantId    uint32 `json:"merchantId" gorm:"not null;"`                  // 商户ID
	CommentatorId uint32 `json:"commentator_id" gorm:"not null;"`              // 评论者ID
	Content       string `json:"content" gorm:"not null;"`                     // 评论内容
}

// TableName /** 复写默认方法，设置表名
func (UserEvaluate) TableName() string {
	return "user_evaluate"
}
