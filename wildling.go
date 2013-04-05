package main

import (
	"flag"
	//Use the relative import path here so project skeleton works ootb. 
	//users should change this to the appropriate absolute path, since relative imports
	//are not part of the golang spec
	"./controllers"
	"github.com/trendrr/cheshire-golang/cheshire"
	"log"
)

func init() {
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
	log.Println("starting app with config=" + config)

	log.Println("Starting")
	//starts listening on all configured interfaces
	bootstrap.Start()
}
