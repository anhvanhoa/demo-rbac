package user

import (
	"iris/services"

	"github.com/kataras/iris/v12"
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
