package main

import (
	"app/config"
	"app/rbac"
	"app/routes"
	"app/services"

	configtech "github.com/TechMaster/core/config"
	"github.com/TechMaster/core/db"
	"github.com/anhvanhoa/lib/middlewares"
	rbacanhvanhoa "github.com/anhvanhoa/lib/rbac"
	routesanhvanhoa "github.com/anhvanhoa/lib/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	app := iris.New()
	// Load config
	configtech.ReadConfig()
	// Init the database
	db.ConnectPostgresqlDB()
	defer db.DB.Close()
	// Load rbac role
	rbac.LoadRole()
	// Config file static
	config.InitFileStatic(app)
	// config Cors
	crs := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
		AllowCredentials: true,
	})
	app.UseRouter(crs)
	// Load rules from db
	routesanhvanhoa.LoadRoutes(
		services.GetRbacRules,
		routes.RulesAuth,
		routes.RulesRbac,
	)
	app.Use(middlewares.RBACMiddleware(&routesanhvanhoa.AllRouter, func(ctx iris.Context) ([]rbacanhvanhoa.Role, error) {
		id := ctx.GetCookie("id")
		user, err := services.GetInforUser(id)
		if err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{
				"message": "Unauthorized",
			})
			return []rbacanhvanhoa.Role{}, err
		}
		return user.Roles, nil
	}, func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.JSON(iris.Map{
			"message": "Forbidden",
		})
	}))
	// Register router
	routes.InitRouters(app)

	app.Run(iris.Addr(config.GetVipper("port")), iris.WithSitemap(config.GetVipper("host")))
}
