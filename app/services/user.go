package services

import (
	"app/models"

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
			users.id, users.username;
	`
	var user = new(models.User)
	_, err := db.DB.QueryOne(user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUser(id string) (*models.ResponseUser, error) {
	query := `
		SELECT 
			users.id, 
			users.username, 
			users.email,
			coalesce(
				array_agg(
					json_build_object(
						'id', roles.id,
						'name', roles.name
					)
				) FILTER (WHERE roles.id IS NOT NULL),
				'{}'
			) AS roles
		FROM 
			users
		LEFT JOIN 
			user_roles ON users.id = user_roles.user_id
		LEFT JOIN 
			roles ON user_roles.role_id = roles.id
		WHERE
			users.id = ?
		GROUP BY 
			users.id;
	`
	var user = new(models.ResponseUser)
	_, err := db.DB.QueryOne(user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAllUser() ([]models.UserAll, error) {
	query := `
		SELECT 
			users.id, 
			users.username, 
			users.email, 
			array_agg(roles.name) AS roles
		FROM 
			users
		LEFT JOIN 
			user_roles ON users.id = user_roles.user_id
		LEFT JOIN 
			roles ON user_roles.role_id = roles.id
		GROUP BY 
			users.id, users.username;
	`
	var users []models.UserAll
	_, err := db.DB.Query(&users, query)
	if err != nil {
		return users, err
	}
	return users, nil
}

func AddRolesToUsers(body models.RequestRolesToUser) error {
	userRoles := []models.UserRole{}

	for _, roleId := range body.RoleIds {
		userRole := models.UserRole{
			UserId: body.UserId,
			RoleId: roleId,
		}
		userRoles = append(userRoles, userRole)
	}
	_, err := db.DB.Model(&userRoles).Insert()
	if err != nil {
		return err
	}
	return nil
}

func DeleteRolesToUsers(body models.RequestRolesToUser) error {
	userRoles := []models.UserRole{}
	for _, roleId := range body.RoleIds {
		userRole := models.UserRole{
			UserId: body.UserId,
			RoleId: roleId,
		}
		userRoles = append(userRoles, userRole)
	}
	_, err := db.DB.Model(&userRoles).Delete()
	if err != nil {
		return err
	}
	return nil
}
