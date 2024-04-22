package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
