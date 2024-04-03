package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
)

// GET /auctions
func FindAuctions(c *gin.Context) {
	var auctions []models.Auction
	models.DB.Find(&auctions)

	c.JSON(http.StatusOK, gin.H{
		"data": auctions,
	})
}
