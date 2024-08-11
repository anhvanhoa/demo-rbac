package auth

import (
	"iris/models"
	"iris/services"
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12"
)

func (a *Auth) Login(ctx iris.Context) {
	var login models.UserLogin
	if err := ctx.ReadJSON(&login); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	user, err := services.Login(login.Email, login.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": "Email or password is incorrect"})
		return
	}
	value := strconv.Itoa(int(user.Id))
	ctx.SetCookie(&http.Cookie{
		Name:  "id",
		Value: value,
		Path:  "/",
	})
	ctx.JSON(*user)
}
func (a *Auth) ViewLogin(ctx iris.Context) {
	ctx.View("pages/login")
}
