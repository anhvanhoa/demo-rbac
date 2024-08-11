package todo

import "github.com/kataras/iris/v12"

type Todo struct{}

type InterfaceTodo interface {
	GetTodos(ctx iris.Context)
}

func NewTodo() InterfaceTodo {
	return &Todo{}
}
