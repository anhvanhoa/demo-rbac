package routes

import (
	"iris/models"
	"log"

	"github.com/anhvanhoa/lib/rbac"
	aroutes "github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type Router struct {
	Path string
	Api  *iris.Application
}

func ConvertRuleFormDb(ruleDb []models.RbacRule) []aroutes.Rule {
	var rules []aroutes.Rule
	for _, rule := range ruleDb {
		var auth rbac.AuthFunc
		switch rule.AuthType {
		case "ALLOW":
			auth = rbac.Allow(rule.Roles...)
		case "DENY":
			auth = rbac.Deny(rule.Roles...)
		case "ALLOW_ALL":
			auth = rbac.AllowAll()
		case "DENY_ALL":
			auth = rbac.DenyAll()
		case "ALLOW_ADMIN":
			auth = rbac.AllowAdmin()
		default:
			log.Printf("Unknown auth type: %s", rule.AuthType)
			continue
		}
		// Add rule to rules
		rules = append(
			rules,
			aroutes.Rule{
				Path:   rule.Path,
				Method: rule.Method,
				Auth:   auth,
				Status: rule.Status,
				Roles:  rule.Roles,
			},
		)
	}
	return rules
}

func InitRouter(app *iris.Application, path string) (*Router, router.Party) {
	return &Router{
		Path: path,
		Api:  app,
	}, app.Party(path)
}
