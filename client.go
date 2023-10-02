package main

import (
	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/ogios/simple-proxy-server/log"
)

func InitClient(conn *normal.Conn) error {
	key, err := conn.Si.GetSec()
	if err != nil {
		return DealError(err, conn)
	}
	strkey := string(key)
	log.Info(nil, "(Client)Get key: %s", strkey)
	data, err := conn.Si.GetSec()
	if err != nil {
		return DealError(err, conn)
	}
	log.Debug(nil, "(Client)Get data: %s", string(data))
	log.Info(nil, "(Client)Start Processing")
	err = Process(conn, strkey, data, TYPE_CLIENT)
	if err != nil {
		log.Info(nil, "(Client)Process done with error: %v", err)
		return err
	}
	log.Info(nil, "(Client)Process done")
	return nil
}
