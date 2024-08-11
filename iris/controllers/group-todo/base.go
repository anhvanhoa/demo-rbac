package grouptodo

import "github.com/kataras/iris/v12"

type GroupTodo struct{}

type InterfaceGroupTodo interface {
	GetGroupTodos(ctx iris.Context)
	GetGroupTodo(ctx iris.Context)
}

func NewTodoGroupController() InterfaceGroupTodo {
	return &GroupTodo{}
}
