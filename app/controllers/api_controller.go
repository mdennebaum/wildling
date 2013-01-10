package controllers

import(
	"github.com/trendrr/cheshire-golang/cheshire"
    "github.com/trendrr/cheshire-golang/strest"    
    "time"
)

// init function. use for registering controllers
func init(){ 
    //register the ping controller.
    cheshire.RegisterApi("/ping", "GET", Ping)
    cheshire.RegisterApi("/firehose", "GET", Firehose)
}

// a demo Ping controller function
func Ping(request *strest.Request,conn strest.Connection) {
        response := strest.NewResponse(request)
        response.Put("data", "PONG")
        conn.Write(response)
}

// a demo Firehose controller
func Firehose(request *strest.Request,conn strest.Connection) {
    for i :=0; true; i++ {
        response := strest.NewResponse(request)
        response.Put("iteration", i)
        response.Put("data", "This is a firehose, I never stop")
        response.SetTxnStatus("continue") //set the status to continue so clients know to expect more responses
        conn.Write(response)

        time.Sleep(200 * time.Millisecond)            
    }
}