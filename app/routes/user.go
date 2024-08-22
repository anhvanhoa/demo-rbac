package routes

import (
	userCtl "app/controllers/user"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	aroutes "github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var userPaths map[string]string = map[string]string{
	"users":              "/users",
	"user":               "/users/{id}",
	"users-add-roles":    "/users/add-roles",
	"users-delete-roles": "/users/delete-roles",
}

func InitRouterUser(app *iris.Application) {
	_, group := InitRouter(app, "/users")
	ctl := userCtl.NewUserController()
	group.Get("/", ctl.GetUsers)
	group.Get("/{id}", ctl.GetUser)
	group.Post("/add-roles", ctl.AddRoles)
	group.Post("/delete-roles", ctl.DeleteRoles)
}

var RulesUser = []aroutes.Rule{
	{
		Path:   userPaths["users"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
	{
		Path:   userPaths["user"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
	{
		Path:   userPaths["users-add-roles"],
		Method: http.MethodPost,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
	{
		Path:   userPaths["users-delete-roles"],
		Method: http.MethodPost,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}
