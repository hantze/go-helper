package main

import (
	"logger/config"
	"logger/internal/httpserver"
)

func main() {
	server := httpserver.NewHTTPServer()
	server.Serve(config.GetConf("HTTP_PORT"))
}

