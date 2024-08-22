package services

import (
	"iris/models"

	"github.com/TechMaster/core/db"
)

func GetInforUser(id string) (*models.User, error) {
	query := `
	SELECT 
		users.id, 
		users.username, 
		users.email, 
		users.password,
		users.avatar,
		array_agg(roles.id) AS roles
	FROM 
		users
	JOIN 
		user_roles ON users.id = user_roles.user_id
	JOIN 
		roles ON user_roles.role_id = roles.id
	WHERE 
		users.id = ?
	GROUP BY 
		users.id;
	`
	var user = new(models.User)
	_, err := db.DB.QueryOne(user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}
