package controllers

import(
	"github.com/trendrr/cheshire-golang/cheshire"
    "github.com/trendrr/cheshire-golang/strest"    
    "log"
)

func init(){ 
    log.Println("init1")

    //register the ping controller.
    cheshire.RegisterApi("/ping", []string{"GET"}, Ping)

	// helloapi := new(HelloAPIController)
	// helloapi.HandleMethod("GET",helloapi.Get())
	// helloapi.Route("/api",helloapi)
}

// type HelloAPIController struct{
// 	cheshire.APIController
// }

// func (this *HelloAPIController)Get()func(){
// 	return func(){
// 		this.Write(this.NewSuccessJson("oh baby baby your a wildling"))
// 	}
// }

// a demo Ping controller
func Ping(request *strest.Request,conn strest.Connection) {
        response := strest.NewResponse(request)
        response.Put("data", "PONG")
        conn.Write(response)
}