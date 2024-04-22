package controllers

import (
	"errors"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadmp97/DomainBids/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func randStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// POST /auth
func Login(c *gin.Context) {
	var userCredentials UserCredentials
	if err := c.ShouldBind(&userCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The request body should contain username/password."})
		return
	}

	var user models.User
	result := models.DB.Where("username = ?", userCredentials.Username).First(&user)

	token := randStringRunes(64)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userCredentials.Password), bcrypt.DefaultCost)
		newUser := models.User{Nickname: "Unknown", Username: userCredentials.Username, Password: string(hashedPassword), Token: token, CreatedAt: time.Now().Format(time.RFC3339)}
		models.DB.Create(&newUser)
		c.JSON(http.StatusCreated, gin.H{"data": gin.H{"token": token}})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found!"})
		return
	}

	user.Token = token
	models.DB.Save(&user)

	c.JSON(http.StatusCreated, gin.H{"data": gin.H{"token": token}})
}

// GET /auth
func CheckAuthentication(c *gin.Context) {
	token := strings.Split(c.GetHeader("Authorization"), " ")

	// No token
	if len(token) == 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"data": gin.H{"status": false}})
		return
	}

	var userCount int64
	models.DB.Model(&models.User{}).Where("token = ?", token[1]).Count(&userCount)
	if userCount == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"data": gin.H{"status": false}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"status": true}})
}
