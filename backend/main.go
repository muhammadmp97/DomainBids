package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadmp97/DomainBids/app/controllers"
	"github.com/muhammadmp97/DomainBids/app/middlewares"
	"github.com/muhammadmp97/DomainBids/app/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDatabase()

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	r.GET("/ping", controllers.Ping)

	r.GET("/auth", controllers.CheckAuthentication)
	r.POST("/auth", controllers.Login)

	r.GET("/users/:id", controllers.FindUser)

	r.GET("/auctions", controllers.FindAuctions)
	r.GET("/auctions/:id", controllers.FindAuction)
	r.POST("/auctions", controllers.StartAuction)

	r.Run("127.0.0.1:" + os.Getenv("APP_PORT"))
}
