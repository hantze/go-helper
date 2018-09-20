package config

import "os"

var defaultConf = map[string]string{
	"HTTP_PORT": "3333",
}

func GetConf(key string) string {
	r := os.Getenv(key)
	if r == "" {
		r = defaultConf[key]
	}
	return r
}
