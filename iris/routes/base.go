package routes

import (
	"github.com/kataras/iris/v12"
)

func InitRouters(app *iris.Application) {
	InitRouterTodo(app)
	InitRoutersUser(app)
	InitRoutersAuth(app)
	InitRoutersSingle(app)
	InitRouterTodoGroups(app)
}
