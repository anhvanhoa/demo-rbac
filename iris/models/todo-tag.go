package models

type TodoTag struct {
	tableName struct{} `pg:"todo_tags"`

	ID     uint `pg:",pk"`
	TodoID uint `pg:",notnull"`
	TagID  uint `pg:",notnull"`
}
