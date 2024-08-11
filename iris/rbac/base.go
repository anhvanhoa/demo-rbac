package rbac

import (
	"iris/services"
	"strings"
)

type RolesType = map[string]int

var Roles RolesType = map[string]int{}

func LoadRole() {
	roles, err := services.GetAllRole()
	if err != nil {
		panic(err)
	}
	for _, role := range roles {
		var name = strings.ToUpper(role.Name)
		Roles[name] = role.Id
	}
}
