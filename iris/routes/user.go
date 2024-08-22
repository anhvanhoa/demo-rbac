package routes

import (
	useCtl "iris/controllers/user"
	"net/http"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var userPaths map[string]string = map[string]string{
	"account": "/account",
}

var RulesUser = []routes.Rule{
	{
		Path:   userPaths["account"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}

func InitRoutersUser(app *iris.Application) {
	_, user := InitRouter(app, "/account")
	ctl := useCtl.NewUserController()
	user.Get("/", ctl.GetUser)
}
