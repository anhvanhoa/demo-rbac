package middlewares

import (
	"fmt"
	"iris/services"
	"iris/tmpl"

	"github.com/kataras/iris/v12"
)

func CheckAuth(ctx iris.Context) {
	data := map[string]interface{}{
		"header": tmpl.HeaderData,
	}
	id := ctx.GetCookie("id")
	user, err := services.GetInforUser(id)
	if err != nil {
		data["isLogin"] = false
	} else {
		data["isLogin"] = user
		fmt.Println(user.Username)
	}
	ctx.ViewData("header", data)
	ctx.Next()
}
