package main

import (
	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/ogios/simple-proxy-server/config"
	"github.com/ogios/simple-proxy-server/log"
)

func main() {
	server, err := normal.NewSocketServer(config.GlobalConfig.Address)
	if err != nil {
		panic(err)
	}
	log.Info("Server created")

	server.AddTypeCallback("client", InitClient)
	server.AddTypeCallback("server", InitServer)
	log.Info("Callback added")

	if err := server.Serv(); err != nil {
		panic(err)
	}
}
