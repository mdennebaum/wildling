package main
import (
	"log"
	"github.com/trendrr/cheshire-golang/cheshire"
	"flag"
	"github.com/mdennebaum/wildling/app/controllers"
)

func init(){
	flag.StringVar(&config, "config", "config/config.yaml", "path to the app config.yaml")
}

var config string 

func main() {

	//parse the command line args
	flag.Parse()

	bootstrap := cheshire.NewBootstrapFile(config)

	//make sure the linker includes our controllers and runs inits
	//this is mandatory
	controllers.Load()  
	
	//tell everyone we started up
	log.Println("starting app with config="+config)

	log.Println("Starting")
    //starts listening on all configured interfaces
    bootstrap.Start()
}
