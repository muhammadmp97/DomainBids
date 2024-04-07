package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"gorm.io/gorm"
)

// GET /auctions
func FindAuctions(c *gin.Context) {
	var auctions []models.Auction
	models.DB.Find(&auctions)

	c.JSON(http.StatusOK, gin.H{
		"data": auctions,
	})
}

func FindAuction(c *gin.Context) {
	var auction models.Auction
	result := models.DB.Preload("Bids.User").First(&auction, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": auction,
	})
}
