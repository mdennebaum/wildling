package controllers

import (
	"github.com/trendrr/goshire/cheshire"
)

func init() {
	cheshire.RegisterHtml("/", "GET", Index)
}

//an example html page
func Index(txn *cheshire.Txn) {
	//create a context map to be passed to the template
	context := make(map[string]interface{})
	context["content"] = "Welcome to the wild(ing)!"

	//set a flash message
	cheshire.Flash(txn, "success", "this is a flash message!")

	//Render index template in layout
	cheshire.RenderInLayout(txn, "/public/index.html", "/layouts/base.html", context)
}
