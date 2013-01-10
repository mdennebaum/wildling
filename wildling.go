package main
import (
	"log"
	"github.com/trendrr/cheshire-golang/cheshire"
	"flag"
	"os"
	"fmt"
)

func init(){
	flag.StringVar(&config, "config", "relative", "path to the app config.yaml")
}

var config string 

func main() {

	//parse the command line args
	flag.Parse()

	//make sure the linker includes our controllers and runs inits
	// controllers.Load()

	//if the config isn't specified on the command line then try the default relative location
	if config == "relative" {
		//get our current working dir
		app_path, _ := os.Getwd()
		//create our config path based on assumption
		config = fmt.Sprintf("%s/config/config.yaml", app_path)
	}
	bootstrap := cheshire.NewBootstrapFile(config)
	log.Println(bootstrap)
	//tell everyone we started up
	log.Println("starting app with config="+config)
}

/*
import (
	"fmt"
	"github.com/trendrr/cheshire-golang"
	"app/controllers"
	"log"
	"os"
	"flag"
)

var config string 

func init(){
	flag.StringVar(&config, "config", "relative", "path to the app config.yaml")
}

func main() {
	//parse the command line args
	flag.Parse()

	//make sure the linker includes our controllers and runs inits
	controllers.Load()

	//if the config isn't specified on the command line then try the default relative location
	if config == "relative" {
		//get our current working dir
		app_path, _ := os.Getwd()
		//create our config path based on assumption
		config = fmt.Sprintf("%s/config/config.yaml", app_path)
	}

	//tell everyone we started up
	log.Println("starting app with config="+config)
	
	//start a new server
	cheshire.NewBootstrap(config).Cheshire.Start()
}
*/