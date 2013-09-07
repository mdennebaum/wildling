package main

import (
	"flag"
	//Use the relative import path here so project skeleton works ootb.
	//users should change this to the appropriate absolute path, since relative imports
	//are not part of the golang spec
	"./controllers"
	"github.com/trendrr/goshire/cheshire"
	"github.com/trendrr/goshire/cheshire/impl/gocache"
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

	//Setup our cache.  this uses the local cache
	//you will need
	//github.com/pmylund/go-cache
	cache := gocache.New(10, 10)
	bootstrap.AddFilters(cheshire.NewSession(cache, 3600))

	//make sure the linker includes our controllers and runs inits
	//this is mandatory
	controllers.Load()

	//tell everyone we started up
	log.Println("starting app with config=" + config)

	//starts listening on all configured interfaces
	bootstrap.Start()
}
