package user

import (
	"iris/services"
	"slices"

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
	roles, err := services.GetAllRole()
	if err != nil {
		panic(err)
	}
	rolesName := []string{}
	for _, role := range roles {
		ok := slices.Contains(result.Roles, role.Id)
		if ok {
			rolesName = append(rolesName, role.Name)
		}
	}
	ctx.ViewData("user", result)
	ctx.ViewData("roles", rolesName)
	ctx.View("pages/account")
}
