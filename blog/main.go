// Package main is the entry point of the application
//
// This file contains the main function.
// The main function is the entry point of the application.
// It is responsible for initializing the database, models, controllers, and server.
// It then starts the server.
//
// For more information on the main function, see:
// https://golang.org/doc/effective_go.html#main
package main

import (
	"log"

	"blog/config"
	"blog/controllers"
	"blog/database"
	"blog/models"
	"blog/server"
)

func main() {
	// Load config
	// The Load function is defined in blog/config/config.go.
	// It returns a pointer to a Config.
	// For more information on the Config struct, see:
	// blog/config/config.go
	cfg := config.Load()

	// Get config values
	port := cfg.ServerPort()
	dburl := cfg.LoadDBUrl()

	// Init database
	// The NewDatabase function is defined in blog/database/database.go.
	// It returns a pointer to a Database.
	// For more information on the Database struct, see:
	// blog/database/database.go
	db := database.NewDatabase()

	// The InitDB function is defined in blog/database/database.go.
	// It takes a database URL as an argument.
	// It returns an error if the database fails to initialize.
	// If the database fails to initialize, we log the error and exit the program.
	if err := db.InitDB(dburl); err != nil {
		log.Fatalf("%v", err)
	}
	defer db.CloseDB() // Close database connection

	// Init models
	// The NewBlogModel function is defined in blog/models/blog.go.
	// It takes a pointer to a Database as an argument.
	// It returns a pointer to a BlogModel.
	// For more information on the BlogModel struct, see:
	// blog/models/blog.go
	blogModel := models.NewBlogModel(db.GetDB())

	// Init controllers
	// The NewBlogController function is defined in blog/controllers/blog.go.
	// It takes a pointer to a BlogModel as an argument.
	// It returns a pointer to a BlogController.
	// For more information on the BlogController struct, see:
	// blog/controllers/blog.go
	blogController := controllers.NewBlogController(blogModel)

	// Init router
	// Create a new router.
	// The NewRouter function is defined in blog/server/router.go.
	// It takes a pointer to a BlogController as an argument.
	// It returns a pointer to a gin.Engine.
	router := server.NewRouter(blogController)

	// Create a new server.
	// The NewServer function is defined in blog/server/server.go.
	// It takes a pointer to a gin.Engine and a port string as arguments.
	// It returns a pointer to a Server.
	// For more information on the Server struct, see:
	// blog/server/server.go
	srv := server.NewServer(router, port)

	// Start server
	srv.Run()
}
