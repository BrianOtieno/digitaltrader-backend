package controllers

import (
	"digitaltrader/auth"
	"digitaltrader/models"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": "invalid json input"})
		return
	}

	if err := user.CreateUser(); err != nil {
		if err.Error() == "user already exists" {
			c.JSON(409, gin.H{"msg": "User already exists"})
		} else {
			c.JSON(500, gin.H{"msg": "Error creating user"})
		}
		return
	}

	c.JSON(201, gin.H{"user": gin.H{
		"id":       user.ID,
		"username": user.Username,
	}})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := models.GetAllUsers(&users); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, &users)
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := models.GetUserByID(&user, id); err != nil {
		c.JSON(409, gin.H{"msg": "User Not Found"})
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := models.DeleteUser(&user, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := models.FindFirstUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "User not in database!"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": "invalid json input"})
		return
	}

	if err := models.UpdateUser(&user, id); err != nil {
		c.AbortWithStatus(409)
		c.JSON(409, gin.H{"msg": "The request encountered a conflict"})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func Login(c *gin.Context) {
	var loginData models.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid JSON input"})
		return
	}

	var user models.User
	if err := models.GetUserByUsername(&user, loginData.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Invalid username or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Invalid username or password",
		})
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET"),
		Issuer:            "AuthService",
		AccessExpiration:  720,    // Token expires in 720 minutes
		RefreshExpiration: 144000, // Refresh token expires in 1 day
	}

	accessToken, refreshToken, err := jwtWrapper.GenerateTokens(user.Username, user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"msg": "Error generating tokens"})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"username":      user.Username,
		"firstname":     user.FirstName,
		"lastname":      user.LastName,
		"profile_photo": user.ProfilePhoto,
		"role":          user.Role, // Include role in response
	})
}

func RefreshToken(c *gin.Context) {
	var tokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
		Username     string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid JSON input"})
		return
	}

	var user models.User
	if err := models.GetUserByUsername(&user, tokenRequest.Username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	valid, err := auth.VerifyRefreshTokenByUsername(tokenRequest.Username, tokenRequest.RefreshToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid or expired refresh token"})
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET"),
		Issuer:            "AuthService",
		AccessExpiration:  30,     // Expire in 30 minutes
		RefreshExpiration: 144000, // Expire in 2400 hours or 100 days
	}
	accessToken, refreshToken, err := jwtWrapper.GenerateTokens(user.Username, user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error generating tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"role":          user.Role, // Include role in response
	})
}
