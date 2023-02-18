package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *BlogControllers) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
