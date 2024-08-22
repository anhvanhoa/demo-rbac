package single

import (
	"app/services"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

func (t *Single) RefreshRules(ctx iris.Context) {
	routes.LoadRoutes(services.GetRbacRules, routes.AllRouter)
	ctx.JSON(iris.Map{
		"message": "Refresh rules admin successfully",
	})
}
func (t *Single) RefreshRoles(ctx iris.Context) {
	rbac.LoadRole(services.GetAllRole)
	ctx.JSON(iris.Map{
		"message": "Refresh role admin successfully",
	})
}
