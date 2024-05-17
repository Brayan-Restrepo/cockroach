package main

import (
	"cockroach/config"
	"cockroach/database"
	"cockroach/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
