package controllers

import(
	"github.com/mdennebaum/cheshire"
)

func init(){
	helloapi := new(HelloAPIController)
	helloapi.HandleMethod("GET",helloapi.Get())
	helloapi.Route("/api",helloapi)
}

type HelloAPIController struct{
	cheshire.APIController
}

func (this *HelloAPIController)Get()func(){
	return func(){
		this.Write(this.NewSuccessJson("oh baby baby your a wildling"))
	}
}