package rulerbac

import "github.com/kataras/iris/v12"

type RuleRbac struct{}

type InterfaceRuleRbac interface {
	GetRules(ctx iris.Context)
}

func NewRuleRbacController() InterfaceRuleRbac {
	return &RuleRbac{}
}
