package grouptodo

import (
	"iris/services"

	"github.com/kataras/iris/v12"
)

func (t *GroupTodo) GetGroupTodos(ctx iris.Context) {
	id := ctx.GetCookie("id")
	result, err := services.GetGroupTodos(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.ViewData("groupTodos", result)
	ctx.View("pages/group-todo")
}

func (t *GroupTodo) GetGroupTodo(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	result, err := services.GetGroupTodoById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.ViewData("groupTodo", result)
	ctx.View("pages/group-todo-detail")
}
