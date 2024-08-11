package main

import (
	"iris/config"
	"iris/constants"
	"iris/middlewares"
	"iris/rbac"
	"iris/routes"
	"iris/tmpl"

	configtech "github.com/TechMaster/core/config"
	"github.com/TechMaster/core/db"
	"github.com/TechMaster/core/template"
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
	// Config the template
	template.InitBlockEngine(app, constants.DIRECTORY, constants.LAYOUT_DEFAULT)
	tmpl.Funcs()
	// Config file static
	config.InitFileStatic(app)
	// config Cors
	crs := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
		AllowCredentials: true,
	})
	app.UseRouter(crs)
	// Check auth
	app.Use(middlewares.CheckAuth)
	// Load rules from db
	// var rulesDB = routes.ConvertRuleFormDb(*services.GetRbacRules())
	// rulesDefault := routes.RoutesDefault(
	// 	routes.RulesTodo,
	// 	routes.RulesUser,
	// 	routes.RulesSingle,
	// 	routes.RulesAuth,
	// 	routes.RulesTodoGroups,
	// )
	// allRoutes := routes.MergerRules(rulesDB, rulesDefault)
	app.Use(middlewares.RBACMiddleware(&(routes.AllRouter)))
	// load mutil language
	config.InitLanguage(app, constants.VI)
	// Register router
	routes.InitRouters(app)

	app.Run(iris.Addr(config.GetVipper("port")), iris.WithSitemap(config.GetVipper("host")))
}
