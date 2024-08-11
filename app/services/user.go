package services

import (
	"app/models"

	"github.com/TechMaster/core/db"
)

func GetInforUser(id string) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	var user = new(models.User)
	_, err := db.DB.QueryOne(user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAllUser() ([]models.User, error) {
	query := "SELECT * FROM users"
	var users []models.User
	_, err := db.DB.Query(&users, query)
	if err != nil {
		return users, err
	}
	return users, nil
}
