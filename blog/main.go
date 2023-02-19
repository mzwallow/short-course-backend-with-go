package main

import (
	"blog/config"
	"blog/server"
)

func main() {
	cfg := config.Load()
	port := cfg.ServerPort()
	// dburl := cfg.LoadDBUrl()

	// db := database.NewDatabase()

	r := server.NewRouter()
	srv := server.NewServer(r, port)
	srv.Run()
}
