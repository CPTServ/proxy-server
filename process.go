package main

import (
	"runtime"
	"time"
)

type ServerInfo struct {
	Timeout time.Time
	Data    string
}

func (s *ServerInfo) RefreshTimeout() {
	s.Timeout = time.Now().Add(TIMEOUT)
}

func (s *ServerInfo) IsTimeout() bool {
	return s.Timeout.Compare(time.Now()) != 1
}

// global session
var (
	SESSION_MAP map[string]*ServerInfo = map[string]*ServerInfo{}
	TIMEOUT                            = time.Second * 30
)

func SetServer(key string, data string) {
	s := &ServerInfo{
		Timeout: time.Now(),
		Data:    data,
	}
	s.RefreshTimeout()
	SESSION_MAP[key] = s
}

func GetServer(key string) (string, bool) {
	a, ok := SESSION_MAP[key]
	if ok {
		if a.IsTimeout() {
			return "", false
		}
		return a.Data, ok
	}
	return "", false
}

func ClearServer() {
	for {
		time.Sleep(time.Second * 10)
		for key, val := range SESSION_MAP {
			if val.IsTimeout() {
				delete(SESSION_MAP, key)
			}
		}
		runtime.GC()
	}
}
