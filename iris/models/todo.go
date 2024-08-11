package models

type Todo struct {
	tableName struct{} `pg:"todos"`

	ID          uint   `pg:",pk"`
	Name        string `pg:",notnull"`
	Description string ``
	UserID      uint   `pg:",notnull"`
}

type TodoItem struct {
	ID          uint
	Name        string
	Description string
	UserID      uint
	TodoGroupID uint
}
