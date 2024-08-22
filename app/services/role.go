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

func DeleteRole(id string) error {
	var query = "DELETE FROM roles WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	return err
}

func SynchronousUser(id string) error {
	var query = "DELETE FROM user_roles WHERE role_id = ?"
	_, err := db.DB.Exec(query, id, id)
	return err
}
func SynchronousRule(id string) error {
	var query = "DELETE FROM rule_roles WHERE role_id = ?"
	_, err := db.DB.Exec(query, id, id)
	return err
}
