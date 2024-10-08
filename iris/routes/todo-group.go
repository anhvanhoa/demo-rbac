package routes

import (
	grouptodo "iris/controllers/group-todo"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var todoGroupPaths map[string]string = map[string]string{
	"todo-groups": "/todo-groups",
	"todo-group":  "/todo-groups/{id}",
}

var RulesTodoGroups = []routes.Rule{
	{
		Path:   todoGroupPaths["todo-groups"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
	{
		Path:   todoGroupPaths["todo-group"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}

func InitRouterTodoGroups(app *iris.Application) {
	_, group := InitRouter(app, "/todo-groups")
	ctl := grouptodo.NewTodoGroupController()
	group.Get("/", ctl.GetGroupTodos)
	group.Get("/{id}", ctl.GetGroupTodo)
}
