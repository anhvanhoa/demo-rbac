package services

import (
	"github.com/TechMaster/core/db"
	"github.com/anhvanhoa/lib/models"
)

func GetAllRole() ([]models.Role, error) {
	var query = "SELECT * FROM roles"
	var roles []models.Role		
	if _, err := db.DB.Query(&roles, query); err != nil {
		return roles, err
	}
	return roles, nil
}
