This is a starter app for the [go cheshire](https://github.com/trendrr/cheshire-golang) web framework. You can use this as a boilerplate when starting a new cheshire project. 

##Install

To get started make sure you have the following installed

```
go get github.com/kylelemons/go-gypsy/yaml
go get github.com/hoisie/mustache
go get code.google.com/p/go.net/websocket
go get github.com/trendrr/cheshire-golang
go get github.com/mdennebaum/wildling
```

##Config

Edit the config/config.yaml file to match with your settings.

##Run

Once you've edited the config.yaml to your liking run the following. 

```
go run wildling.go -config=/path/to/config.yaml
```

Alternatively if your config folder is located in the same dir as the executable you can omit the -config flag

```
go run wildling.go
```

Then you should be able to point your browser at http://localhost:8000/ and http://localhost:8000/api to see examples of html controller output and api controller output. 

##More

For more info on how to use the go cheshire framework checkout https://github.com/trendrr/cheshire-golang. If there are any issues please submit an ticket. 

