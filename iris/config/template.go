package config

import "github.com/kataras/iris/v12/view"

func InitTemplate(dir string, ext string, layout string) *view.BlocksEngine {
	tmpl := view.Blocks(dir, ext)
	tmpl.Layout(layout)
	tmpl.Load()
	return tmpl
}
