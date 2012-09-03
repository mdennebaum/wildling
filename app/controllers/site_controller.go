package controllers

import(
	"github.com/mdennebaum/cheshire"
)

func init(){
	hellohtml := new(HelloHTMLController)
	hellohtml.HandleMethod("GET",hellohtml.Get())
	hellohtml.Route("/", hellohtml)
}

type HelloHTMLController struct{
	cheshire.HTMLController
}

func (this *HelloHTMLController)Get()func(){
	return func(){
		this.Write("oh baby baby your a wildling...")
	}
}