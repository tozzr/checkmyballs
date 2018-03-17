package actions

import (
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
			"rating": func(r int) template.HTML {
				rating := ""
				for i := 0; i < 5; i++ {
					a := "-o"
					if i < r {
						a = ""
					}
					rating += "<i class='fa fa-star" + a + "'></i>"
				}
				return template.HTML(rating)
			},
		},
	})
}
