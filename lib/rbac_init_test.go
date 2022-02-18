package lib

import (
	"fmt"
	"testing"
)

func init() {
	InitGormDB()
}

func TestGetRoles(t *testing.T) {
	m := make([]*RoleRel, 0)
	GetRoles(4, &m)
	for _, rel := range m {
		fmt.Println(rel.PRole + ":" + rel.Role)
	}
}

func TestGetUserRoles(t *testing.T) {
	userRoles := GetUserRoles()
	for _, role := range userRoles {
		fmt.Println(role.UserName + ":" + role.RoleName)
	}
}

func TestGetRoleRoutes(t *testing.T) {
	roleRoutes := GetRoleRoutes()
	for _, role := range roleRoutes {
		fmt.Println(role.RouteUri + ":" + role.RoleName)
	}
}
