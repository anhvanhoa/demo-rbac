package routes

import (
	"app/controllers/single"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var singlePaths map[string]string = map[string]string{
	"refresh-rules": "/refresh-rules",
	"refresh-roles": "/refresh-roles",
}

var RulesSingle = []routes.Rule{
	{Path: singlePaths["refresh-rules"], Method: http.MethodPost, Auth: rbac.AllowAdmin(), Status: true},
	{Path: singlePaths["refresh-roles"], Method: http.MethodPost, Auth: rbac.AllowAdmin(), Status: true},
}

func InitRoutersSingle(app *iris.Application) {
	_, group := InitRouter(app, "")
	ctl := single.NewSingleController()
	group.Post("/refresh-rules", ctl.RefreshRules)
	group.Post("/refresh-roles", ctl.RefreshRoles)
}
