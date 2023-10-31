package main

import (
	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/ogios/simple-proxy-server/log"
)

func InitClient(conn *normal.Conn) error {
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Error("[Client]%+v", err)
		}
	}()
	key, err := conn.Si.GetSec()
	if err != nil {
		panic(err)
	}
	strkey := string(key)
	log.Info("[Client]Get key: %s", strkey)
	// data, err := conn.Si.GetSec()
	// if err != nil {
	// 	panic(err)
	// }
	// log.Debug("[Client]Get data: %s", string(data))
	log.Debug("[Client]Start Processing")
	// err = Process(conn, strkey, data, TYPE_CLIENT)
	// if err != nil {
	// 	panic(err)
	// }
	addrs, ok := GetServer(string(key))
	if ok {
		err = conn.So.AddBytes([]byte("success"))
		if err != nil {
			panic(err)
		}
		err = conn.So.AddBytes([]byte(addrs))
		if err != nil {
			panic(err)
		}
	} else {
		err = conn.So.AddBytes([]byte("error"))
		if err != nil {
			panic(err)
		}
		err = conn.So.AddBytes([]byte("No addrs found"))
		if err != nil {
			panic(err)
		}
	}
	err = conn.So.WriteTo(conn.Raw)
	if err != nil {
		panic(err)
	}
	log.Debug("[Client]Process done")
	return nil
}
