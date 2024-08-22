package services

import (
	"fmt"

	"github.com/TechMaster/core/db"
	"github.com/anhvanhoa/lib/models"
)

func GetRbacRules() []models.RbacRule {
	query := `
		SELECT 
			rbac_rules.id,
			rbac_rules.path,
			rbac_rules.method,
			rbac_rules.auth_type,
			rbac_rules.name,
			rbac_rules.status,
			rbac_rules.service,
			array_agg(roles.id) AS roles
		FROM 
			rbac_rules
		LEFT JOIN 
			rule_roles ON rbac_rules.id = rule_roles.rule_id
		LEFT JOIN 
			roles ON rule_roles.role_id = roles.id
		WHERE 
			service = 'MAIN'
		GROUP BY 
			rbac_rules.id;
	`
	var rules *[]models.RbacRule = new([]models.RbacRule)
	if _, err := db.DB.Query(rules, query); err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}
	return *rules
}
