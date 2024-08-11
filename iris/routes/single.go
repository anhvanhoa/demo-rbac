package routes

import (
	// "iris/controllers/rule"
	// "iris/config"
	"iris/controllers/single"
	"iris/rbac"
	"net/http"

	"github.com/kataras/iris/v12"
)

var singlePaths map[string]string = map[string]string{
	"home":    "/",
	"refresh": "/refresh-rules",
}

var RulesSingle = []Rule{
	{Path: singlePaths["home"], Method: http.MethodGet, Auth: rbac.AllowAll(), Status: false},
	{Path: singlePaths["refresh"], Method: http.MethodPost, Auth: rbac.AllowAdmin(), Status: true},
}

func InitRoutersSingle(app *iris.Application) {
	_, group := InitRouter(app, "")
	ctl := single.NewSingleController()
	// ctlRule := rule.NewRuleController()
	group.Get("/", ctl.Home)
	group.Post("/refresh-rules", func(ctx iris.Context) {
		// ctl.RefreshRules(ctx)
		// RulesDB = ConvertRuleFormDb(*services.GetRbacRules())
		LoadRoutes()
		ctx.JSON(iris.Map{
			"message": "Refresh rules successfully",
		})
	})
}
