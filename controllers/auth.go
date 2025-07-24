package controllers

import (
	"net/http"
	"time"

	"github.com/HrushikeshAnandSarangi/go-rest/config"
	"github.com/HrushikeshAnandSarangi/go-rest/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtkey = []byte("secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Health Check": "OK"})
}
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	input.Password = string(hashedPassword)
	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"message": "Registration Successful"})

}
func Login(c *gin.Context) {
	var input models.User
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Where("email = ?", input.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: input.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtkey)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
