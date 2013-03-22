package controllers

import(
    "github.com/trendrr/cheshire-golang/cheshire"   
    "time"
)

// init function. use for registering controllers
func init(){ 
    //register the ping controller.
    cheshire.RegisterApi("/ping", "GET", Ping)
    cheshire.RegisterApi("/firehose", "GET", Firehose)
}

// a demo Ping controller function
func Ping(request *cheshire.Request,conn cheshire.Connection) {
        response := cheshire.NewResponse()
        response.Put("data", "PONG")
        conn.Write(response)
}

// a demo Firehose controller
func Firehose(request *cheshire.Request,conn cheshire.Connection) {
    for i :=0; true; i++ {
        response := cheshire.NewResponse()
        response.Put("iteration", i)
        response.Put("data", "This is a firehose, I never stop")
        response.SetTxnStatus("continue") //set the status to continue so clients know to expect more responses
        conn.Write(response)

        time.Sleep(200 * time.Millisecond)            
    }
}
