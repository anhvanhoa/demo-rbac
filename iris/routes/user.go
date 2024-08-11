package routes

import (
	userCtl "iris/controllers/user"
	"iris/rbac"
	"net/http"

	"github.com/kataras/iris/v12"
)

var userPaths map[string]string = map[string]string{
	"account": "/account",
}

var RulesUser = []Rule{
	{
		Path:   userPaths["account"],
		Method: http.MethodGet,
		Auth:   rbac.AllowAdmin(),
		Status: true,
	},
}

func InitRoutersUser(app *iris.Application) {
	_, user := InitRouter(app, "/account")
	ctl := userCtl.NewUserController()
	user.Get("/", ctl.GetUser)
}
