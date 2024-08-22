package rulerbac

import (
	"app/services"

	"github.com/kataras/iris/v12"
)

func (t *RuleRbac) GetRules(ctx iris.Context) {

	result := services.GetAllRules()
	// if err != nil {
	// 	ctx.StatusCode(iris.StatusInternalServerError)
	// 	ctx.JSON(iris.Map{"message": err.Error()})
	// 	return
	// }
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(result)
}
