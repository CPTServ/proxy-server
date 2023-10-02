package main

import (
	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/ogios/simple-proxy-server/log"
)

func InitServer(conn *normal.Conn) error {
	key, err := conn.Si.GetSec()
	if err != nil {
		return DealError(err, conn)
	}
	strkey := string(key)
	log.Info(nil, "(Server)Get key: %s", strkey)
	data, err := conn.Si.GetSec()
	if err != nil {
		return DealError(err, conn)
	}
	log.Debug(nil, "(Server)Get data: %s", string(data))
	log.Info(nil, "(Server)Start Processing")
	err = Process(conn, strkey, data, TYPE_SERVER)
	if err != nil {
		log.Info(nil, "(Server)Process done with error: %v", err)
		return err
	}
	log.Info(nil, "(Server)Process done")
	return nil
}
