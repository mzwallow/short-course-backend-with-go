package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
}

func NewBlogController() *BlogController {
	return &BlogController{}
}

func (c *BlogController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
