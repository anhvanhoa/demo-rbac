package routes

import (
	"github.com/kataras/iris/v12"
)

func InitRouters(app *iris.Application) {
	InitRoutersAuth(app)
	InitRouterRules(app)
	InitRouterRole(app)
	InitRoutersSingle(app)
	InitRouterUser(app)
}
