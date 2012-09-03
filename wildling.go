package main

import(
	"os"
	"fmt"
	"log"
	"github.com/mdennebaum/cheshire"
	"wildling/app/controllers"
)

func main(){
	log.Println("starting wildling...")
	
	//get our current working dir
	app_path,_ := os.Getwd()
	
	//make sure we load our controllers
	controllers.Load()
	
	//start a new server
	cheshire.NewBootstrap(fmt.Sprintf("%s/config/config.yaml", app_path)).Cheshire.Start()
}
