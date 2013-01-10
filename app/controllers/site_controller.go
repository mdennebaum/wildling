package controllers

import(
	// "github.com/trendrr/cheshire-golang/cheshire"
	"log"
)


func init(){
    log.Println("init2")

	// helloapi := new(HelloAPIController)
	// helloapi.HandleMethod("GET",helloapi.Get())
	// helloapi.Route("/api",helloapi)
}
// func init(){
// 	hellohtml := new(HelloHTMLController)
// 	hellohtml.Layout = "/layouts/base.html"
// 	hellohtml.HandleMethod("GET",hellohtml.Get())
// 	hellohtml.Route("/", hellohtml)
// }

// type HelloHTMLController struct{
// 	cheshire.HTMLController
// }

// func (this *HelloHTMLController)Get()func(){
// 	return func() {
// 		this.ContextVar("test", "a test view var")
// 		this.RenderTemplate("/public/index.html")
// 	}
// }