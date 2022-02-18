package model

import (
	"time"
)

// Role 权限表
type Role struct {
	RoleID      int       `gorm:"column:role_id;primary_key;AUTO_INCREMENT"`             // 权限id
	RoleName    string    `gorm:"column:role_name;NOT NULL"`                             // 权限名
	RolePid     int       `gorm:"column:role_pid;NOT NULL"`                              // 父权限id
	RoleComment string    `gorm:"column:role_comment;NOT NULL"`                          // 权限备注
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *Role) TableName() string {
	return "t_role"
}
