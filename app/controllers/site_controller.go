package controllers

import(
	"github.com/mdennebaum/cheshire"
)

func init(){
	hellohtml := new(HelloHTMLController)
	hellohtml.Layout = "/layouts/base.html"
	hellohtml.HandleMethod("GET",hellohtml.Get())
	hellohtml.Route("/", hellohtml)
}

type HelloHTMLController struct{
	cheshire.HTMLController
}

func (this *HelloHTMLController)Get()func(){
	return func() {
		this.ContextVar("test", "a test view var")
		this.RenderTemplate("/public/index.html")
	}
}