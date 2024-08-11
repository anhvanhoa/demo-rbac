package services

import (
	"fmt"
	"iris/models"

	"github.com/TechMaster/core/db"
)

func GetGroupTodos(id string) ([]models.GroupTodo, error) {
	query := `SELECT * FROM group_todos WHERE user_id = ?`
	var groupTodos []models.GroupTodo
	_, err := db.DB.Query(&groupTodos, query, id)
	if err != nil {
		return groupTodos, err
	}
	return groupTodos, nil
}

func GetGroupTodoById(id int) ([]models.GroupTodoByUser, error) {
	query := `SELECT todos.id, todos.name, todos.description FROM group_todos join todos ON group_todos.id = todos.todo_group_id WHERE group_todos.id = ?`
	var groupTodo []models.GroupTodoByUser
	_, err := db.DB.Query(&groupTodo, query, id)
	fmt.Println(groupTodo)
	if err != nil {
		return groupTodo, err
	}
	return groupTodo, nil
}
