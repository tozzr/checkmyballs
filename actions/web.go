package actions

import (
	"html/template"

	"github.com/gobuffalo/buffalo"
)

func SetCurrentUrl(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("current_url", c.Request().URL.Path)
		return next(c)
	}
}

func RatingRenderHelper(r int) template.HTML {
	rating := ""
	for i := 0; i < 5; i++ {
		a := "-o"
		if i < r {
			a = ""
		}
		rating += "<i class='fa fa-star" + a + "'></i>"
	}
	return template.HTML(rating)
}

func ActiceClassRenderHelper(path string, url string) string {
	if path == url {
		return "active"
	}
	return ""
}
