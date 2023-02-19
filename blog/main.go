package main

import (
	"blog/config"
	"fmt"
)

func main() {
	cfg := config.Load()
	port := cfg.ServerPort()
	dburl := cfg.LoadDBUrl()

	fmt.Printf("port: %v\n", port)
	fmt.Printf("dburl: %v\n", dburl)
}
