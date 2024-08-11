package routes

import (
	"iris/controllers/todo"
	"iris/rbac"
	"net/http"

	"github.com/kataras/iris/v12"
)

var todoPaths map[string]string = map[string]string{
	"todos": "/todos",
	"todo":  "/todos/{id}",
}

var RulesTodo = []Rule{
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
