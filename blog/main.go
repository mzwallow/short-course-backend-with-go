package main

import (
	"log"
	"os"

	"blog/config"
	"blog/controllers"
	"blog/database"
	"blog/models"
	"blog/server"
)

func main() {
	cfg := config.NewConfig()
	cfg.Load()

	db := database.NewDatabase()
	if err := db.Connect(cfg.GetDatabaseURL()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	blogModels := models.NewBlogModels(db.DB())

	blogControllers := controllers.NewBlogControllers(blogModels)

	router := server.NewRouter(blogControllers)

	srv := server.NewServer(router, cfg.GetPort())
	srv.Run()
}
