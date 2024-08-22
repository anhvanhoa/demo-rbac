package routes

import (
	"app/controllers/role"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	aroutes "github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var rolePaths map[string]string = map[string]string{
	"roles":        "/roles",
	"delete-roles": "/roles/delete/{id:int}",
}

func InitRouterRole(app *iris.Application) {
	_, group := InitRouter(app, "/roles")
	ctl := role.NewRoleController()
	group.Get("/", ctl.GetRoles)
	group.Delete("/delete/{id:int}", ctl.DeleteRole)
}

var RulesRole = []aroutes.Rule{
	{
		Path:   rolePaths["roles"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
	{
		Path:   rolePaths["delete-roles"],
		Method: http.MethodDelete,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}
