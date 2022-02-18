package lib

import (
	"log"
	"permissions-learning/model"
)

type RoleRel struct {
	PRole string
	Role  string
}

func GetRoles(pid int, m *[]*RoleRel) {
	roles := make([]*model.Role, 0)
	rolesTree := make(map[int][]*model.Role, 0)
	rolesMap := make(map[int]*model.Role, 0)
	rolesMap[0] = &model.Role{
		RoleID:   0,
		RoleName: "super_admin",
	}

	err := Gorm.Table("t_role").Find(&roles).Error
	if err != nil {
		log.Fatalln(err)
	}
	for _, role := range roles {
		rolesTree[role.RolePid] = append(rolesTree[role.RolePid], role)
		rolesMap[role.RoleID] = role
	}

	var travel func(id int)
	travel = func(id int) {
		role, ok := rolesTree[id]
		if !ok {
			return
		}

		for _, r := range role {
			*m = append(*m, &RoleRel{
				PRole: rolesMap[r.RolePid].RoleName,
				Role:  r.RoleName,
			})
			travel(r.RoleID)
		}
	}

	travel(pid)
}

func GetUserRoles() []*model.User {
	users := make([]*model.User, 0)
	err := Gorm.Select("user_name, role_name").
		Table("t_user u, t_user_role ur, t_role r").
		Where("u.user_id = ur.user_id and ur.role_id = r.role_id").
		Order("u.user_id desc").
		Find(&users).
		Error
	if err != nil {
		log.Fatalln(err)
	}
	return users
}

func GetRoleRoutes() []*model.Route {
	routes := make([]*model.Route, 0)
	err := Gorm.Select("rou.route_name, rou.route_uri, rou.route_method, rol.role_name").
		Table("t_role rol,t_role_route rr,t_route rou").
		Where("rol.role_id = rr.role_id and rr.route_id = rou.route_id").
		Order("rou.route_id desc").
		Find(&routes).
		Error
	if err != nil {
		log.Fatalln(err)
	}
	return routes
}
