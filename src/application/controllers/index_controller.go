package controllers

import(
	"github.com/mdennebaum/cheshire.go/cheshire/router"
	"github.com/mdennebaum/cheshire.go/cheshire/controller"
)

func init(){
	index := new(IndexController)
	index.HTMLController.HandleMethod("GET",index.Get())
	
	//make sure this path is authenticated. 
	//pick up the global auth mechanism configured in our bootstrap
	//index.BaseController.AuthRequired()
	
	//or you could use the role required to require auth with a specific role
	//index.BaseController.RoleRequired("admin")
	
	router.Instance().AddRoute("/",index)
}

type IndexController struct{
	controller.HTMLController
}

func (this *IndexController)Get()func(){
	return func(){
		this.Write("oh baby baby your a wildling")
	}
}