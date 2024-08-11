package rule

import (
	// "iris/config"
	// "iris/services"

	"github.com/kataras/iris/v12"
)



func (s *Rule) RefreshRules(ctx iris.Context) {
	// config.RulesDB = config.ConvertRuleFormDb(*services.GetRbacRules())
	ctx.JSON(iris.Map{
		"message": "Refresh rules successfully",
	})
}
