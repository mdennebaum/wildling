package wildling

import(
	"log"
	"wildling/db/mongodb"
	"github.com/mdennebaum/cheshire.go/cheshire/bootstrap"
)

type WildlingBootstrap struct{
	App *bootstrap.CheshireBootstrap
}

func (this *WildlingBootstrap)InitMongo(){
	servers := this.App.Conf.Get("mongo.servers")
	if servers != "" {
		//init the mongo connection
		mongodb.Instance().Connect(servers)
		
		//set the default database
		mongodb.Instance().SetDB(this.App.Conf.Get("mongo.name"))
	}else{
		log.Panicln("Can't connect to mongo. Exiting...")
	}
	
}

func NewWildlingBootstrap(configPath string) *WildlingBootstrap{
	bstrap := &WildlingBootstrap{App:bootstrap.NewCheshireBootstrap(configPath)}
	bstrap.InitMongo()
	return bstrap
}