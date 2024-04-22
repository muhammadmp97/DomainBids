package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"github.com/muhammadmp97/DomainBids/app/utils"
)

func init() {
	go runTicker()
}

var channel chan bool

var domainsQueue []string
var domainsPrices map[string]uint = make(map[string]uint)

type HumbleWorthValuation struct {
	Domain string `json:"domain"`
	Price  uint   `json:"auction"`
}

type HumbleWorthResponse struct {
	Valuations []HumbleWorthValuation `json:"valuations"`
}

func runTicker() {
	channel = make(chan bool)
	ticker := time.Tick(3 * time.Second)
	for range ticker {
		if len(domainsQueue) > 0 {
			requestBody, _ := json.Marshal(map[string][]string{"domains": domainsQueue})
			req, _ := http.NewRequest(http.MethodPost, "https://humbleworth.com/api/valuation", bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, _ := client.Do(req)
			defer response.Body.Close()

			body, _ := ioutil.ReadAll(response.Body)

			var jsonResponse HumbleWorthResponse
			json.Unmarshal(body, &jsonResponse)

			for _, valuation := range jsonResponse.Valuations {
				domainsPrices[valuation.Domain] = valuation.Price
			}

			clear(domainsQueue)

			channel <- true
		}
	}
}

func EstimatePrice(c *gin.Context) {
	var user *models.User = utils.GetUserFromToken(c)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	domain, err := url.Parse(c.Param("domain"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Domain is not valid!"})
	}

	if !slices.Contains(domainsQueue, domain.Path) {
		domainsQueue = append(domainsQueue, domain.Path)
	}

	for {
		select {
		case <-channel:
			_, ready := domainsPrices[domain.Path]
			if ready {
				c.JSON(http.StatusOK, gin.H{"data": gin.H{"price": domainsPrices[domain.Path]}})
				delete(domainsPrices, domain.Path)
				return
			}
		}
	}
}
