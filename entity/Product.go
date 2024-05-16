// Package entity /** 商品表
package entity

type Product struct {
	ID             uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 主键 商品ID
	Name           string `json:"name" gorm:"not null;"`                        // 商品名字
	TypeId         uint32 `json:"type_id" gorm:"not null;"`                     // 商品类型ID
	UserUId        string `json:"user_uid" gorm:"not null;"`                    // 商品用户UID
	Icon           string `json:"icon" gorm:"not null;"`                        // 商品图标
	Price          uint32 `json:"price" gorm:"not null;"`                       // 商品价格
	Content        string `json:"content" gorm:"not null;"`                     // 商品介绍
	Recommend      uint8  `json:"recommend" gorm:"not null;"`                   // 推荐
	NotRecommended uint8  `json:"not_recommended" gorm:"not null;"`             // 不推荐
	Collect        uint32 `json:"collect" gorm:"not null default:0;"`           // 收藏
	Status         bool   `json:"status" gorm:"not null; default:1;"`           // 商品是否上架
	UID            string `json:"uid" gorm:"not null;"`                         // 商品标识
}

// TableName /** 复写默认方法，设置表名
func (Product) TableName() string {
	return "product"
}
