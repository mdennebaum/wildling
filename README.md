This is a starter app for the go cheshire web framework. You can use this as a boilerplate when starting a new cheshire project. 

##Install

To get started make sure you have the following installed

```
go get github.com/bradfitz/gomemcache/memcache
go get github.com/kylelemons/go-gypsy/yaml
go get github.com/hoisie/mustache
go get labix.org/v2/mgo
go get github.com/mdennebaum/cheshire
```

##Config

Edit the config/config.yaml file to match with your settings. Please note: the config folder must be in the same directory as your executable. In the future you will be able to pass in a config file as a flag when starting the app but for now we will just rely on convention. 

##Run

Once you've edited the config.yaml to your liking run the following. 

```
go run wildling.go
``` 

Then you should be able to point your browser at http://localhost:8000/ and http://localhost:8000/api to see examples of html controller output and api controller output. 

##More

For more info on how to use the go cheshire framework checkout http://github.com/mdennebaum/cheshire. If there are any issues please submit an ticket. 

