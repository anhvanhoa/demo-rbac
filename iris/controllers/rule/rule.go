package rule

import (
	"iris/services"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

func (s *Rule) RefreshRules(ctx iris.Context) {
	routes.LoadRoutes(services.GetRbacRules, routes.AllRouter)
	ctx.JSON(iris.Map{
		"message": "Refresh rules successfully",
	})
}

func (s *Rule) RefreshRoles(ctx iris.Context) {
	rbac.LoadRole(services.GetAllRole)
	ctx.JSON(iris.Map{
		"message": "Refresh role successfully",
	})
}
