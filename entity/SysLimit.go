// Package entity /** 系统用户类型
package entity

type SysLimit struct {
	ID   uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null;"` // 主键 权限ID
	Type string `json:"type" gorm:"not null;"`                        // 类型
}

// TableName /** 复写默认方法，设置表名
func (SysLimit) TableName() string {
	return "sys_limit"
}
