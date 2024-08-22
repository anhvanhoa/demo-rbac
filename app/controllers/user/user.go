package user

import (
	"app/models"
	"app/services"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/kataras/iris/v12"
)

func (t *User) GetUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := services.GetUser(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	var roles []int
	for _, role := range result.Roles {
		roles = append(roles, role.Id)
	}
	rules, err := services.GetRulesAccess(roles)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.JSON(iris.Map{
		"user":  result,
		"rules": rules,
	})
}

func (t *User) GetUsers(ctx iris.Context) {
	result, err := services.GetAllUser()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.JSON(result)
}
func (t *User) AddRoles(ctx iris.Context) {
	var requetsData models.RequestRolesToUser
	err := ctx.ReadJSON(&requetsData)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	err = services.AddRolesToUsers(requetsData)
	if err != nil {
		ctx.JSON(iris.Map{"message": "1 User đã có role này, không thểm thêm role này vào user"})
		return
	}
	ctx.JSON(iris.Map{"message": "Add roles to users successfully"})
}

func (t *User) DeleteRoles(ctx iris.Context) {
	var requetsData models.RequestRolesToUser
	err := ctx.ReadJSON(&requetsData)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if rbac.AllowAdmin()(requetsData.RoleIds) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Không thể xóa role admin"})
		return
	}
	err = services.DeleteRolesToUsers(requetsData)
	if err != nil {
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"message": "Remove roles to users successfully"})
}
