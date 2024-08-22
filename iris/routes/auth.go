package routes

import (
	"iris/controllers/auth"
	"net/http"

	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

var authPaths map[string]string = map[string]string{
	"login":  "/auth/login",
	"logout": "/auth/logout",
}

var RulesAuth = []routes.Rule{
	{Path: authPaths["login"], Method: http.MethodPost, Status: false},
	{Path: authPaths["login"], Method: http.MethodGet, Status: false},
	{Path: authPaths["logout"], Method: http.MethodGet, Status: false},
}

func InitRoutersAuth(app *iris.Application) {
	_, group := InitRouter(app, "/auth")
	ctl := auth.NewAuthController()
	group.Post("/login", ctl.Login)
	group.Get("/login", ctl.ViewLogin)
	group.Get("/logout", ctl.Logout)
}
