package user

import "github.com/kataras/iris/v12"

type User struct{}

type InterfaceUser interface {
	GetUser(ctx iris.Context)
	GetUsers(ctx iris.Context)
	AddRoles(ctx iris.Context)
	DeleteRoles(ctx iris.Context)
}

func NewUserController() InterfaceUser {
	return &User{}
}
