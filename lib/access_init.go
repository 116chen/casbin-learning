package lib

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func initPolicy() {

	// 初始化继承关系
	m := make([]*RoleRel, 0)
	GetRoles(0, &m)
	for _, r := range m {
		_, err := E.AddRoleForUser(r.PRole, r.Role)
		if err != nil {
			log.Fatalln("初始化继承关系失败！")
		}
	}

	// 初始化用户角色
	userRoles := GetUserRoles()
	for _, userRole := range userRoles {
		_, err := E.AddRoleForUser(userRole.UserName, userRole.RoleName)
		if err != nil {
			log.Fatalln("初始化用户角色失败！")
		}
	}

	// 初始化角色路由
	roleRoutes := GetRoleRoutes()
	for _, roleRoute := range roleRoutes {
		_, err := E.AddPolicy(roleRoute.RoleName, roleRoute.RouteUri, roleRoute.RouteMethod)
		if err != nil {
			log.Fatalln("初始化角色路由！")
		}
	}
}

func InitAccess() {
	InitGormDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatalln(err)
	}
	e, err := casbin.NewEnforcer("resource/basis_model.conf", adapter)
	if err != nil {
		log.Fatalln(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatalln(err)
	}
	E = e
	initPolicy()
	InitFunctions()
}
