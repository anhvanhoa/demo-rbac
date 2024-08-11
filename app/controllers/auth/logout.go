package auth

import "github.com/kataras/iris/v12"

func (a *Auth) Logout(ctx iris.Context) {
	ctx.RemoveCookie("id")
	ctx.Redirect("/")
}
