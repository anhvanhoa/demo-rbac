package single

import "github.com/kataras/iris/v12"

type Single struct{}

type InterfaceSingle interface {
	Home(ctx iris.Context)
}

func NewSingleController() InterfaceSingle {
	return &Single{}
}
