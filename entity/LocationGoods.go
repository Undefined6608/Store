// Package entity /** 货物地址表
package entity

type LocationGoods struct {
	ID           uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 主键 商品路径ID
	OrderId      uint32 `json:"order_id" gorm:"not null;"`                   // 订单ID
	Status       uint8  `json:"status" gorm:"not null;"`                     // 订单状态
	PositionLoad string `json:"position" gorm:"not null;type:varchar(255)"`  // 订单地址
}

// TableName /** 复写默认方法，设置表名
func (LocationGoods) TableName() string {
	return "location_goods"
}
