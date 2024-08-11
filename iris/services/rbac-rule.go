package services

import (
	"iris/models"

	"github.com/TechMaster/core/db"
)

func GetRbacRules() *[]models.RbacRule {
	query := "SELECT * FROM rbac_rules WHERE service = 'MAIN'"
	var rules *[]models.RbacRule = new([]models.RbacRule)
	if _, err := db.DB.Query(rules, query); err != nil {
		panic(err)
	}
	return rules
}
