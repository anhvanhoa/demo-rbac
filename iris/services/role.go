package services

import (
	"iris/models"

	"github.com/TechMaster/core/db"
)

func GetAllRole() ([]models.RoleItem, error) {
	var query = "SELECT * FROM roles"
	var roles []models.RoleItem
	if _, err := db.DB.Query(&roles, query); err != nil {
		return roles, err
	}
	return roles, nil
}
