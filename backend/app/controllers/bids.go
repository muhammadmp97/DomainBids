package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"github.com/muhammadmp97/DomainBids/app/utils"
	"gorm.io/gorm"
)

type BidRequest struct {
	Price uint `json:"price" binding:"required,numeric"`
}

// POST /auctions/:id/bids
func PlaceBid(c *gin.Context) {
	var user *models.User = utils.GetUserFromToken(c)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var bidRequest BidRequest
	if err := c.ShouldBind(&bidRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plase enter a valid number."})
		return
	}

	var auction models.Auction
	result := models.DB.Preload("Bids").First(&auction, c.Param("id"))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found!"})
		return
	}

	if (len(auction.Bids) > 0 && bidRequest.Price <= auction.Bids[len(auction.Bids)-1].Price) || bidRequest.Price <= auction.StartingPrice {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "You should enter a higher number!"})
		return
	}

	newBid := models.Bid{UserID: user.ID, AuctionID: auction.ID, Price: bidRequest.Price, CreatedAt: time.Now().Format(time.RFC3339)}
	models.DB.Create(&newBid)

	c.JSON(http.StatusCreated, gin.H{})
}
