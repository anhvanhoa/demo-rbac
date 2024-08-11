package auth

import "github.com/kataras/iris/v12"

type Auth struct{}

type InterfaceAuth interface {
	Login(ctx iris.Context)
	Logout(ctx iris.Context)
}

func NewAuthController() InterfaceAuth {
	return &Auth{}
}
