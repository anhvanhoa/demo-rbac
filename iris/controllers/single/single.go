package single

import (
	"github.com/kataras/iris/v12"
)

func (s *Single) Home(ctx iris.Context) {
	lang := ctx.URLParamDefault("lang", "vi-VN")
	ctx.SetLanguage(lang)
	title := ctx.Tr("Title")
	ctx.Header("Content-Type", "text/html")
	ctx.ViewData("Title", title)
	err := ctx.View("pages/home")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
