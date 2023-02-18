package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port   string
}

func NewServer(r *gin.Engine, port string) *Server {
	return &Server{router: r, port: port}
}

func (s *Server) Run() {
	log.Printf("Server is running on port %s", s.port)
	if err := s.router.Run(s.port); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
