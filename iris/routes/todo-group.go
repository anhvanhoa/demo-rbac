package routes

import (
	grouptodo "iris/controllers/group-todo"
	"iris/rbac"
	"net/http"

	"github.com/kataras/iris/v12"
)

var todoGroupPaths map[string]string = map[string]string{
	"todo-groups": "/todo-groups",
	"todo-group":  "/todo-groups/{id}",
}

var RulesTodoGroups = []Rule{
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
