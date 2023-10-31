package main

// connection types
var (
	TYPE_CLIENT uint8 = 1
	TYPE_SERVER uint8 = 2
)

// global session
var (
	SESSION_MAP map[string]string
)

func init() {
	SESSION_MAP = make(map[string]string)
}

func SetServer(key string, data string) {
	SESSION_MAP[key] = data
}

func GetServer(key string) (string, bool) {
	a, ok := SESSION_MAP[key]
	return a, ok
}
