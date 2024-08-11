package services

import (
	"iris/models"

	"github.com/TechMaster/core/db"
)

func Login(email, password string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = ? AND password = ?"
	var user = new(models.User)
	_, err := db.DB.QueryOne(user, query, email, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
