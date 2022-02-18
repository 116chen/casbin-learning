package model

import (
	"time"
)

// User 用户表
type User struct {
	UserID     int       `gorm:"column:user_id;primary_key;AUTO_INCREMENT"`             // 用户id
	UserName   string    `gorm:"column:user_name;NOT NULL"`                             // 用户名
	RoleName   string    `gorm:"column:role_name;NOT NULL"`                             // 权限名
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *User) TableName() string {
	return "t_user"
}
