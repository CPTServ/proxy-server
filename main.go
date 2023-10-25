package main

import (
	"github.com/ogios/simple-socket-server/server/normal"
	"golang.org/x/exp/slog"

	"github.com/ogios/simple-proxy-server/config"
	"github.com/ogios/simple-proxy-server/log"
)

func main() {
	log.SetLevel(slog.LevelDebug)
	server, err := normal.NewSocketServer(config.GLOBAL_CONFIG.Addr)
	if err != nil {
		panic(err)
	}
	log.Info(nil, "Server created")

	server.AddTypeCallback("client", InitClient)
	server.AddTypeCallback("server", InitServer)
	log.Info(nil, "Callback added")

	if err := server.Serv(); err != nil {
		panic(err)
	}
}
