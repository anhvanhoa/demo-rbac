package todo

import (
	"iris/services"

	"github.com/kataras/iris/v12"
)

func (t *Todo) GetTodos(ctx iris.Context) {
	result, err := services.GetTodos()
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err.Error()})
		return
	}
	ctx.JSON(result)
}
