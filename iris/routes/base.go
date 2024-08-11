package routes

import (
	"iris/services"

	"github.com/kataras/iris/v12"
)

func InitRouters(app *iris.Application) {
	InitRouterTodo(app)
	InitRoutersUser(app)
	InitRoutersSingle(app)
	InitRoutersAuth(app)
	InitRouterTodoGroups(app)
}

var AllRouter []Rule

func LoadRoutes() {
	var rulesDB = ConvertRuleFormDb(*services.GetRbacRules())
	rulesDefault := RoutesDefault(
		RulesTodo,
		RulesUser,
		RulesSingle,
		RulesAuth,
		RulesTodoGroups,
	)
	AllRouter = *MergerRules(rulesDB, rulesDefault)
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
