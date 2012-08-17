package mongodb

import (
	"launchpad.net/mgo"
)

type mongoDB struct{
	Session *mgo.Session
	DB *mgo.Database
}

//singleton instance
var instance *mongoDB

//interface to our memcache singleton
func Instance() (*mongoDB){
	if (instance == nil) {
		instance = new(mongoDB)	
	}
	return instance
}

func (this *mongoDB) Connect(server string){
	var err error
	
	this.Session, err = mgo.Dial(server)
	
	if err != nil {
		panic(err)
	}
	
	defer this.Session.Close()
}

func (this *mongoDB) SetDB(name string){
	this.DB = this.Session.DB(name)
}

func (this *mongoDB) GetCollection(name string) *mgo.Collection{
	return this.DB.C(name)
}
