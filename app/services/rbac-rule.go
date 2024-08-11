package services

import (
	"github.com/TechMaster/core/db"
	"github.com/anhvanhoa/lib/models"
)

func GetRbacRules() []models.RbacRule {
	query := "SELECT * FROM rbac_rules"
	var rules []models.RbacRule
	if _, err := db.DB.Query(&rules, query); err != nil {
		panic(err)
	}
	return rules
}
