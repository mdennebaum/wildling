package controllers

import (
	"github.com/trendrr/cheshire-golang/cheshire"
)

func init() {
	cheshire.RegisterHtml("/", "GET", Index)
}

//an example html page
func Index(request *cheshire.Request, conn *cheshire.HtmlConnection) {
	//create a context map to be passed to the template
	context := make(map[string]interface{})
	context["content"] = "Welcome to the wild(ing)!"

	//TODO: need to use index.html
	conn.Render("/layouts/base.html", context)
}
