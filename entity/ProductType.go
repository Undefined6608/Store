// Package entity /** 商品类型表
package entity

type ProductType struct {
	ID   uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 商品类型ID
	Root uint32 `json:"root" gorm:"not null;default:0"`               // 类型根节点
	Type string `json:"type" gorm:"not null;"`                        // 类型名称
}

// TableName /** 复写默认方法，设置表名
func (ProductType) TableName() string {
	return "product_type"
}
