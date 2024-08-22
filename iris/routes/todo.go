package routes

import (
	"iris/controllers/todo"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var todoPaths map[string]string = map[string]string{
	"todos": "/todos",
	"todo":  "/todos/{id}",
}

var RulesTodo = []routes.Rule{
	{
		Path:   todoPaths["todos"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
	{
		Path:   todoPaths["todo"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}

func InitRouterTodo(app *iris.Application) {
	_, group := InitRouter(app, "/todos")
	ctl := todo.NewTodo()
	group.Get("/{id}", ctl.GetTodos)
}
