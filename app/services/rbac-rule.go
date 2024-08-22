package services

import (
	appModels "app/models"

	"github.com/TechMaster/core/db"
	"github.com/anhvanhoa/lib/models"
	"github.com/go-pg/pg/v10"
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
			service = 'MANAGER'
		GROUP BY 
			rbac_rules.id;
	`
	var rules []models.RbacRule
	if _, err := db.DB.Query(&rules, query); err != nil {
		panic(err)
	}
	return rules
}

func GetAllRules() []appModels.RbacRule {
	query := `
		SELECT 
			rbac_rules.*,
			array_agg(roles.name) AS roles
		FROM 
			rbac_rules
		LEFT JOIN 
			rule_roles ON rbac_rules.id = rule_roles.rule_id
		LEFT JOIN 
			roles ON rule_roles.role_id = roles.id
		GROUP BY 
			rbac_rules.id;
	`
	var rules []appModels.RbacRule
	if _, err := db.DB.Query(&rules, query); err != nil {
		panic(err)
	}
	return rules
}

func GetRulesAccess(role []int) ([]appModels.RbacRule, error) {
	query := `
		SELECT
			rbac_rules.*
		FROM
			rbac_rules
		LEFT JOIN
			rule_roles ON rbac_rules.id = rule_roles.rule_id
		LEFT JOIN
			roles ON rule_roles.role_id = roles.id
		WHERE
			rbac_rules.auth_type IN ('ALLOW_ALL', 'ALLOW_ADMIN')
		    OR
		    roles.id IN (?) AND rbac_rules.auth_type = 'ALLOW'
		GROUP BY
			rbac_rules.id;
	`
	var rules []appModels.RbacRule
	if _, err := db.DB.Query(&rules, query, pg.In(role)); err != nil {
		return nil, err
	}
	return rules, nil
}
