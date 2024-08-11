package routes

import (
	"github.com/kataras/iris/v12"
)

func InitRouters(app *iris.Application) {
	InitRoutersAuth(app)
	InitRouterRules(app)
}

func MergerRules(rulesDb []Rule, rulesDefault []Rule) *[]Rule {
	var rules []Rule
	for _, rule := range rulesDefault {
		var found bool
		for _, ruleDb := range rulesDb {
			if rule.Path == ruleDb.Path && rule.Method == ruleDb.Method {
				rules = append(rules, ruleDb)
				found = true
				break
			}
		}
		if !found {
			rules = append(rules, rule)
		}
	}
	return &rules
}

func RoutesDefault(rules ...[]Rule) []Rule {
	rulesMerge := []Rule{}
	for _, rule := range rules {
		rulesMerge = append(rulesMerge, rule...)
	}
	return rulesMerge
}
