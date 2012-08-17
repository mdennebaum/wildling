package controllers

import(
	"github.com/mdennebaum/cheshire.go/cheshire/router"
	"github.com/mdennebaum/cheshire.go/cheshire/controller"
)

func init(){
	test := new(IndexController2)
	test.HTMLController.HandleMethod("GET",test.Get())
	
	//make sure this path is authenticated. 
	//pick up the global auth mechanism configured in our bootstrap
	//index.BaseController.AuthRequired()
	
	//or you could use the role required to require auth with a specific role
	//index.BaseController.RoleRequired("admin")
	
	router.Instance().AddRoute("/test",test)
}

type IndexController2 struct{
	controller.HTMLController
}

func (this *IndexController2)Get()func(){
	return func(){
		this.Write("ohai werldz test")
	}
}