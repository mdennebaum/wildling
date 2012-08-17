package main


import(
	"os"
	"fmt"
	"log"
	"wildling"
	"application/controllers"
)

func main(){
	log.Println("starting wildling server...")
	
	//get our current working dir
	app_path,_ := os.Getwd()
	
	//make sure we load our controllers
	controllers.Load()
	
	//start a new server
	wildling.NewWildlingBootstrap(fmt.Sprintf("%s/config/config.yaml",app_path)).App.Server.Start()
}