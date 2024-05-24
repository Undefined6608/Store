// Package entity /** 商品轮播图表
/**
 * @projectName:    Store
 * @package:        entity
 * @className:      ProductBanner
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/10 12:42
 * @version:    1.0
 */
package entity

type ProductBanner struct {
	ID          uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 主键轮播图ID
	Banner      string `json:"banner" gorm:"not null;"`                      // 轮播图
	ProductName string `json:"productName" gorm:"not null;"`                 // 商品名称
	ProductID   uint32 `json:"product_id" gorm:"not null;"`                  // 商品ID
	UID         string `json:"uid" gorm:"not null;"`                         // 轮播图标识
}

// TableName /** 复写默认方法，设置表名
func (ProductBanner) TableName() string {
	return "product_banner"
}
