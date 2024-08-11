package routes

import (
	rulerbac "app/controllers/rule"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	aroutes "github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var rulesRbacPaths map[string]string = map[string]string{
	"rules": "/rules",
}

var RulesRbac = []aroutes.Rule{
	{
		Path:   rulesRbacPaths["rules"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}

func InitRouterRules(app *iris.Application) {
	_, group := InitRouter(app, "/rules")
	ctl := rulerbac.NewRuleRbacController()
	group.Get("/", ctl.GetRules)
}
