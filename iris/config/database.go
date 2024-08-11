package config

import (
	"context"
	"fmt"
	"iris/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func InitDatabasePostgreSQL() {
	// Init the database
	db := pg.Connect(&pg.Options{
		Addr:     GetEnv("DB_ADDR"),
		User:     GetEnv("DB_USER"),
		Password: GetEnv("DB_PASSWORD"),
		Database: GetEnv("DB_NAME"),
	})

	DB = db

	LoadModel(db)

	// Log
	if IsModeDev() {
		db.AddQueryHook(dbLogger{})
	}

	// Check the connection
	ctx := db.Context()
	err := db.Ping(ctx)
	if err != nil {
		panic(err)
	}
}

func LoadModel(db *pg.DB) {
	models := []interface{}{
		(*models.UserModel)(nil),
		(*models.Todo)(nil),
		(*models.Tag)(nil),
		(*models.TodoTag)(nil),
		(*models.GroupTodo)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		}
	}
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	fmt.Println("After query :" + string(bytes))
	return nil
}
