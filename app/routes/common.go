package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type Router struct {
	Path string
	Api  *iris.Application
}

func InitRouter(app *iris.Application, path string) (*Router, router.Party) {
	return &Router{
		Path: path,
		Api:  app,
	}, app.Party(path)
}
