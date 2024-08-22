package main

import (
	"fmt"
	"iris/config"
	"iris/constants"
	"iris/middlewares"
	"iris/routes"
	"iris/services"
	"iris/tmpl"
	"regexp"
	"sort"
	"strings"

	configtech "github.com/TechMaster/core/config"
	"github.com/TechMaster/core/db"
	"github.com/TechMaster/core/template"

	// amiddlewares "github.com/anhvanhoa/lib/middlewares"
	"github.com/anhvanhoa/lib/rbac"
	aRouters "github.com/anhvanhoa/lib/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func matchPath(path, rulePath string) bool {
	if path == rulePath {
		return true
	}
	// Chuyển đổi rulePath thành regex
	regexPattern := "^" + regexp.QuoteMeta(rulePath) + "$"
	// Tìm tất cả các tham số động trong rulePath và thay thế chúng bằng regex
	re := regexp.MustCompile(`\\\{[^/]+\\\}`)
	regexPattern = re.ReplaceAllString(regexPattern, `[^/]+`)
	// Kiểm tra xem path có khớp với regex không
	matched, _ := regexp.MatchString(regexPattern, path)
	return matched
}

func sortRules(rules *[]aRouters.Rule) {
	sort.SliceStable(*rules, func(i, j int) bool {
		countParams := func(path string) int {
			return strings.Count(path, "{")
		}
		paramsI := countParams((*rules)[i].Path)
		paramsJ := countParams((*rules)[j].Path)
		if paramsI != paramsJ {
			return paramsI < paramsJ
		}
		return len((*rules)[i].Path) < len((*rules)[j].Path)
	})
}

func RBACMiddleware(rules *[]aRouters.Rule, auth func(ctx iris.Context) ([]rbac.Role, error), handleErrForbidden func(ctx iris.Context)) iris.Handler {
	return func(ctx iris.Context) {
		fmt.Println(ctx.GetCurrentRoute())
		for _, rule := range *rules {
			if matchPath(ctx.Path(), rule.Path) && ctx.Method() == rule.Method {
				if !rule.Status {
					ctx.Next()
					return
				}
				roles, err := auth(ctx)
				if err != nil {
					return
				}
				if rbac.AllowAdmin()(roles) || rule.Auth(roles) {
					ctx.Next()
					return
				}
				handleErrForbidden(ctx)
				return
			}
		}
		ctx.Next()
	}
}

func main() {
	app := iris.New()
	// Load config
	configtech.ReadConfig()
	// Init the database
	db.ConnectPostgresqlDB()
	defer db.DB.Close()
	// Load rbac role
	rbac.LoadRole(services.GetAllRole)
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
	aRouters.LoadRoutes(
		services.GetRbacRules,
		routes.RulesTodo,
		routes.RulesUser,
		routes.RulesAuth,
		routes.RulesSingle,
		routes.RulesTodoGroups,
	)
	sortRules(&aRouters.AllRouter)
	app.Use(RBACMiddleware(
		&aRouters.AllRouter,
		func(ctx iris.Context) ([]rbac.Role, error) {
			id := ctx.GetCookie("id")
			user, err := services.GetInforUser(id)
			if err != nil {
				ctx.StatusCode(iris.StatusUnauthorized)
				ctx.JSON(iris.Map{
					"message": "Unauthorized",
				})
				return []rbac.Role{}, err
			}
			return user.Roles, nil
		}, func(ctx iris.Context) {
			ctx.StatusCode(iris.StatusForbidden)
			ctx.JSON(iris.Map{
				"message": "Forbidden access",
			})
		}))
	// load mutil language
	config.InitLanguage(app, constants.VI)
	// Register router
	routes.InitRouters(app)

	app.Run(iris.Addr(config.GetVipper("port")), iris.WithSitemap(config.GetVipper("host")))
}
