package controllers

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"github.com/muhammadmp97/DomainBids/app/utils"
	"gorm.io/gorm"
)

// GET /auctions
func FindAuctions(c *gin.Context) {
	var auctions []models.Auction
	models.DB.Find(&auctions)

	c.JSON(http.StatusOK, gin.H{"data": auctions})
}

// GET /auctions/:id
func FindAuction(c *gin.Context) {
	var auction models.Auction
	result := models.DB.Preload("Bids.User").First(&auction, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": auction})
}

type AuctionRequest struct {
	SLD           string `json:"sld" binding:"required"`
	TLD           string `json:"tld" binding:"required"`
	StartingPrice uint   `json:"starting_price"`
	Description   string `json:"description"`
}

// POST /auctions
func StartAuction(c *gin.Context) {
	var auctionRequest AuctionRequest
	if err := c.ShouldBind(&auctionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The request is invalid."})
		return
	}

	var user *models.User = utils.GetUserFromToken(c)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	records, _ := net.LookupTXT(fmt.Sprintf("%s.%s", auctionRequest.SLD, auctionRequest.TLD))
	recordWasFound := false
	key := string(strings.Split(c.GetHeader("Authorization"), " ")[1][0:7]) // First seven characters of the token
	for _, record := range records {
		if record == key {
			recordWasFound = true
		}
	}

	if recordWasFound == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "The TXT record was not found!"})
		return
	}

	if auctionRequest.StartingPrice < 200 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Starting price can't be less than $200!"})
		return
	}

	endsAt := time.Now().Add(48 * time.Hour)
	newAuction := models.Auction{SLD: auctionRequest.SLD, TLD: auctionRequest.TLD, Description: auctionRequest.Description, StartingPrice: auctionRequest.StartingPrice, UserID: user.ID, Status: "cr", EndsAt: endsAt.Format(time.RFC3339), CreatedAt: time.Now().Format(time.RFC3339)}
	models.DB.Create(&newAuction)

	c.JSON(http.StatusCreated, gin.H{"data": newAuction})
}
