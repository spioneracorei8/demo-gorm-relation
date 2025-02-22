package main

import (
	"golang_relation/config"
	"golang_relation/server"
)

func getMainServer() *server.Server {
	return &server.Server{
		PSQL_CONNECTION: config.PSQL_CONNECTION,
		APP_PORT:        config.APP_PORT,
	}
}

func main() {
	s := getMainServer()
	s.Start()
}
