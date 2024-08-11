package middlewares

import (
	"app/rbac"
	"app/routes"
	"app/services"
	"regexp"
	"strings"

	"github.com/kataras/iris/v12"
)

func RBACMiddleware(rules *[]routes.Rule) iris.Handler {
	return func(ctx iris.Context) {
		for _, rule := range *rules {
			pass := matchPath(ctx.Path(), rule.Path)
			if pass && ctx.Method() == rule.Method && rule.Status {
				id := ctx.GetCookie("id")
				user, err := services.GetInforUser(id)
				if err != nil {
					ctx.StatusCode(iris.StatusUnauthorized)
					ctx.JSON(iris.Map{"error": "Unauthorized"})
					return
				}
				if rbac.AllowAdmin()(user.Roles) {
					break
				}
				if !rule.Auth(user.Roles) {
					ctx.StatusCode(iris.StatusForbidden)
					ctx.JSON(iris.Map{"error": "Access denied"})
					return
				}
				break
			}
		}
		ctx.Next()
	}
}

func matchPath(path, rulePath string) bool {
	// Chuyển đổi rulePath thành regex
	regexPattern := "^" + regexp.QuoteMeta(rulePath) + "$"
	regexPattern = strings.ReplaceAll(regexPattern, `\{id\}`, `[^/]+`)
	// regexPattern = strings.ReplaceAll(regexPattern, `\{token\}`, `[^/]+`)

	// Kiểm tra xem path có khớp với regex không
	matched, _ := regexp.MatchString(regexPattern, path)
	return matched
}

// func findLongestSubstring(x string, rules []string) (bool, *string) {
// 	// Sắp xếp rules theo thứ tự giảm dần của độ dài
// 	sort.Slice(rules, func(i, j int) bool {
// 		return len(rules[i]) > len(rules[j])
// 	})

// 	// Duyệt qua danh sách rules
// 	for _, rule := range rules {
// 		if strings.Contains(x, rule) {
// 			return true, &rule
// 		}
// 	}

// 	// Nếu không tìm thấy, trả về nil
