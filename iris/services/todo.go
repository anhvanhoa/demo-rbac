package services

import (
	"iris/models"

	"github.com/TechMaster/core/db"
)

func GetTodos() ([]models.TodoItem, error) {
	var todos []models.TodoItem
	query := "SELECT * FROM todos"
	if _, err := db.DB.Query(&todos, query); err != nil {
		return todos, err
	}
	return todos, nil
}
