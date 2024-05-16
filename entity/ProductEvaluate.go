// Package entity /** 商品评论
package entity

type ProductEvaluate struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 商品评论ID
	ProductId uint32 `json:"product_id" gorm:"not null;"`                  // 商品ID
	UserId    uint32 `json:"user_id" gorm:"not null;"`                     // 用户ID
	Content   string `json:"content" gorm:"not null;"`                     // 评论内容
	Icon      string `json:"icon"`                                         // 评论图片
}

// TableName /** 复写默认方法，设置表名
func (ProductEvaluate) TableName() string {
	return "product_evaluate"
}
