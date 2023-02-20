// Package server provides a wrapper around the gin server
//
// This file contains the server struct and the Run method.
//
// The Run method is used to start the server.
package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Server is a wrapper around the gin server.
// It contains the router and the port to run the server on.
// The Run method is used to start the server.
type Server struct {
	// The gin.Engine is the router.
	// For more information on the gin.Engine, see:
	// https://godoc.org/github.com/gin-gonic/gin#Engine
	router *gin.Engine

	// The port to run the server on.
	// It is a string because it can contain a colon and a port number.
	// For example, ":8080".
	port string
}

// NewServer creates a new server.
// It takes a pointer to a gin.Engine and a port string as arguments.
// It returns a pointer to a Server.
func NewServer(r *gin.Engine, port string) *Server {
	return &Server{
		router: r,
		port:   port,
	}
}

// Run starts the server.
func (s *Server) Run() {
	// s.router.Run starts the server.
	// It takes a port string as an argument.
	// It returns an error if the server fails to start.
	// If the server fails to start, we log the error and exit the program.
	//
	// For more information on s.router.Run, see:
	// https://godoc.org/github.com/gin-gonic/gin#Engine.Run
	if err := s.router.Run(s.port); err != nil {
		log.Fatalf("failed to run server on port %s with %v", s.port, err)
	}
}
