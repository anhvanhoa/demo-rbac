package user

import (
	"app/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func (t *User) GetUser(ctx iris.Context) {
	id := ctx.GetCookie("id")
	result, err := services.GetInforUser(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.ViewData("user", result)
	ctx.View("pages/account")
}

func (t *User) GetUsers(ctx iris.Context) {
	result, err := services.GetAllUser()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.JSON(result, context.JSON{Indent: "  ", Secure: true})
}
