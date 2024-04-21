package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"gorm.io/gorm"
)

// GET /users/:id
func FindUser(c *gin.Context) {
	var user models.User
	result := models.DB.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
