package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"gorm.io/gorm"
)

func GetUserFromToken(c *gin.Context) *models.User {
	token := strings.Split(c.GetHeader("Authorization"), " ")

	// No token
	if len(token) == 1 {
		return nil
	}

	var user models.User
	result := models.DB.Where("token = ?", token[1]).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}
