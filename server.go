package main

import (
	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/ogios/simple-proxy-server/log"
)

func InitServer(conn *normal.Conn) error {
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Error("[Server]%+v", err)
		}
	}()
	key, err := conn.Si.GetSec()
	if err != nil {
		panic(err)
	}
	strkey := string(key)
	log.Info("[Server]Get key: %s", strkey)
	data, err := conn.Si.GetSec()
	if err != nil {
		panic(err)
	}
	log.Debug("[Server]Get data: %s", string(data))
	log.Debug("[Server]Start Processing")
	SetServer(string(key), string(data))
	// err = Process(conn, strkey, data, TYPE_SERVER)
	// if err != nil {
	// 	log.Info("[Server]Process done with error: %v", err)
	// 	return err
	// }
	_, err = conn.Raw.Write([]byte{200})
	if err != nil {
		panic(err)
	}
	log.Debug("[Server]Process done")
	return nil
}
