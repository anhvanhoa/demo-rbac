package role

import "github.com/kataras/iris/v12"

type Role struct{}

type InterfaceRole interface {
	GetRoles(ctx iris.Context)
	DeleteRole(ctx iris.Context)
}

func NewRoleController() InterfaceRole {
	return &Role{}
}
