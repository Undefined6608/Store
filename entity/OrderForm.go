// Package entity /** 订单表
package entity

type OrderForm struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 主键 订单ID
	ProductId uint32 `json:"product_id" gorm:"not null"`                  // 商品ID
	UserId    uint32 `json:"user_id" gorm:"not null"`                     // 用户ID
	ToAddress uint32 `json:"to_address" gorm:"not null"`                  // 送货地址ID
	Remarks   string `json:"remarks" gorm:"not null"`                     // 订单备注
	Status    uint8  `json:"status" gorm:"not null;default:1"`            // 发货状态
}

// TableName /** 复写默认方法，设置表名
func (OrderForm) TableName() string {
	return "order_form"
}
