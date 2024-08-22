package role

import (
	"app/services"

	"github.com/kataras/iris/v12"
)

func (t *Role) GetRoles(ctx iris.Context) {
	result, err := services.GetAllRole()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(result)
}

func (t *Role) DeleteRole(ctx iris.Context) {
	id := ctx.Params().Get("id")
	if id == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Missing id"})
		return
	}
	if err := services.SynchronousUser(id); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := services.SynchronousRule(id); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := services.DeleteRole(id); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Delete role successfully"})
}
