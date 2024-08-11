package config

import (
	"app/constants"

	"github.com/kataras/iris/v12"
)

var optionStatic = iris.DirOptions{
	Compress: true,
}

func InitFileStatic(app *iris.Application) {
	app.HandleDir(constants.REQUEST_PATH_STATIC, constants.FOLDER_STATIC, optionStatic)
}
