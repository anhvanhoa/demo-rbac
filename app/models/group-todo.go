package models

type GroupTodo struct {
	tableName struct{} `pg:"group_todos"`

	ID     uint   `pg:",pk"`
	Name   string `pg:",notnull"`
	UserID uint   `pg:",notnull"`
}

type GroupTodoByUser struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
