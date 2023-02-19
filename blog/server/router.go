package server

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(blogCtrl *controllers.BlogController) *gin.Engine {
	r := gin.Default()

	r.GET("/health", blogCtrl.Health)

	return r
}
