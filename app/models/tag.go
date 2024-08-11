package models 

type Tag struct {
	tableName struct{} `pg:"tags"`

	ID   uint   `pg:",pk"`
	Name string `pg:",notnull"`
	Color string `pg:",notnull"`
}