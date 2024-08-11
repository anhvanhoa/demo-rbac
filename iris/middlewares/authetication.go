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
	_, err := services.GetInforUser(id)
	if err != nil {
		data["isLogin"] = false
		fmt.Println("CheckAuth: ", "false")
	} else {
		data["isLogin"] = "true"
		fmt.Println("CheckAuth: ", true)
	}
	ctx.ViewData("header", data)
	ctx.Next()
}
