package rule

import "github.com/kataras/iris/v12"

type Rule struct{}

type InterfaceRule interface {
	RefreshRules(ctx iris.Context)
	RefreshRoles(ctx iris.Context)
}

func NewRuleController() InterfaceRule {
	return &Rule{}
}
