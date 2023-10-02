package main

import (
	"context"
	"sync"
	"time"

	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/ogios/simple-proxy-server/log"
)

// address with connection
type Addrs struct {
	Conn   *normal.Conn
	Data   []byte
	Origin string
}

// both server and client conn&addr&key&context
type ConnInfo struct {
	Server Addrs
	Client Addrs
	Cond   sync.Cond
	Key    string
	Ctx    context.Context
	Cancel context.CancelFunc
}

// connection types
var TYPE_CLIENT uint8 = 1
var TYPE_SERVER uint8 = 2

// global session
var SESSION_MAP map[string]*ConnInfo
var MAP_COND sync.Cond

func init() {
	SESSION_MAP = make(map[string]*ConnInfo)
	MAP_COND = *sync.NewCond(&sync.Mutex{})
}

// process error and send it to the socket connetion
func DealError(err error, conn *normal.Conn) error {
	defer conn.Close()
	se := err.Error()
	err = conn.So.AddBytes([]byte("error"))
	if err == nil {
		err = conn.So.AddBytes([]byte(se))
		if err == nil {
			err = conn.So.WriteTo(conn.Raw)
		}
	}
	return err
}

// process a connection, if the other one is connected already then send addresses to each other
func Process(conn *normal.Conn, strkey string, data []byte, t uint8) error {
	MAP_COND.L.Lock()
	defer MAP_COND.L.Unlock()
	log.Debug(nil, "prcess %d", t)
	i, ok := SESSION_MAP[strkey]
	if !ok {
		log.Debug(nil, "No session registered, create one")
		i = new(ConnInfo)
		i.Cond = *sync.NewCond(&sync.Mutex{})
		i.Key = strkey
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		i.Cancel = cancel
		i.Ctx = ctx
	}
	log.Debug(nil, "creating address")
	addr := &Addrs{
		Origin: conn.Raw.RemoteAddr().String(),
		Data:   data,
		Conn:   conn,
	}

	log.Debug(nil, "adding address")
	var check Addrs
	switch t {
	case TYPE_CLIENT:
		i.Client = *addr
		check = i.Server
	case TYPE_SERVER:
		i.Server = *addr
		check = i.Client
	}
	SESSION_MAP[strkey] = i

	if check.Origin != "" {
		log.Debug(nil, "start communicating...")
		err := i.Conmunicate()
		if err != nil {
			return err
		}
	}
	return nil
}

// send addresses to each other
func (c *ConnInfo) Conmunicate() error {
	defer delete(SESSION_MAP, c.Key)
	defer c.Cancel()
	defer c.Client.Conn.Close()
	defer c.Server.Conn.Close()
	log.Debug(nil, "adding client data...")
	err := c.Client.Conn.So.AddBytes(c.Client.Data)
	if err != nil {
		return err
	}
	err = c.Client.Conn.So.AddBytes([]byte(c.Client.Origin))
	if err != nil {
		return err
	}
	log.Debug(nil, "sending to server...")
	err = c.Client.Conn.So.WriteTo(c.Server.Conn.Raw)
	if err != nil {
		return err
	}
	log.Debug(nil, "adding server data...")
	err = c.Server.Conn.So.AddBytes(c.Server.Data)
	if err != nil {
		return err
	}
	err = c.Server.Conn.So.AddBytes([]byte(c.Server.Origin))
	if err != nil {
		return err
	}
	log.Debug(nil, "sending to client...")
	err = c.Server.Conn.So.WriteTo(c.Client.Conn.Raw)
	if err != nil {
		return err
	}
	return nil
}
