package controllers

import (
	"github.com/trendrr/goshire/cheshire"
	"time"
)

// init function. use for registering controllers
func init() {
	//register the ping controller.
	cheshire.RegisterApi("/ping", "GET", Ping)
	cheshire.RegisterApi("/firehose", "GET", Firehose)
}

// a demo Ping controller function
func Ping(txn *cheshire.Txn) {
	response := cheshire.NewResponse(txn)
	response.Put("data", "PONG")
	txn.Write(response)
}

// a demo Firehose controller
func Firehose(txn *cheshire.Txn) {
	for i := 0; true; i++ {
		response := cheshire.NewResponse(txn)
		response.Put("iteration", i)
		response.Put("data", "This is a firehose, I never stop")
		response.SetTxnStatus("continue") //set the status to continue so clients know to expect more responses
		txn.Write(response)

		time.Sleep(200 * time.Millisecond)
	}
}
