package config

import (
	"iris/constants"

	"github.com/kataras/iris/v12"
)

func InitLanguage(app *iris.Application, langDefault string) {
	err := app.I18n.Load(constants.PATH_FOLDER_LANG, constants.EN, constants.VI)
	app.I18n.SetDefault(langDefault)
	if err != nil {
		panic(err)
	}
}
