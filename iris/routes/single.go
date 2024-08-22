package routes

import (
	"iris/controllers/rule"
	"iris/controllers/single"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var singlePaths map[string]string = map[string]string{
	"home":          "/",
	"refresh-rules": "/refresh-rules",
	"refresh-roles": "/refresh-roles",
	"test":          "/{id:int}/{token}",
}

var RulesSingle = []routes.Rule{
	{Path: singlePaths["home"], Method: http.MethodGet, Auth: rbac.AllowAll(), Status: false},
	{Path: singlePaths["test"], Method: http.MethodGet, Auth: rbac.AllowAll(), Status: true},
	{Path: singlePaths["refresh-rules"], Method: http.MethodPost, Auth: rbac.AllowAdmin(), Status: true},
	{Path: singlePaths["refresh-roles"], Method: http.MethodPost, Auth: rbac.AllowAdmin(), Status: true},
}

func InitRoutersSingle(app *iris.Application) {
	_, group := InitRouter(app, "")
	ctl := single.NewSingleController()
	ctlRule := rule.NewRuleController()
	group.Get("/", ctl.Home)
	group.Get("/{id:int}/{token}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		token := ctx.Params().Get("token")
		ctx.WriteString("ID: " + id + " - Token: " + token)
	})
	group.Post("/refresh-rules", ctlRule.RefreshRules)
	group.Post("/refresh-roles", ctlRule.RefreshRoles)
}
