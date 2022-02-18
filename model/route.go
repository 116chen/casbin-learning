package model

import (
	"time"
)

// Route 路由表
type Route struct {
	RouteID     int       `gorm:"column:route_id;primary_key;AUTO_INCREMENT"`            // 路由id
	RouteName   string    `gorm:"column:route_name;NOT NULL"`                            // 路由名
	RoleName    string    `gorm:"column:role_name;NOT NULL"`                             // 权限名
	RoutePid    int       `gorm:"column:route_pid;NOT NULL"`                             // 父路由id
	RouteUri    string    `gorm:"column:route_uri;NOT NULL"`                             // 路由uri
	RouteMethod string    `gorm:"column:route_method;NOT NULL"`                          // 方法
	RouteState  string    `gorm:"column:route_state;NOT NULL"`                           // 状态
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *Route) TableName() string {
	return "t_route"
}
