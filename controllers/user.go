// controllers/user_controller.go
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

	// Bind incoming JSON request to the User model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": "invalid json input"})
		return
	}

	// Call the CreateUser method from the model to create the user
	if err := user.CreateUser(); err != nil {
		if err.Error() == "user already exists" {
			c.JSON(409, gin.H{"msg": "User already exists"})
		} else {
			c.JSON(500, gin.H{"msg": "Error creating user"})
		}
		return
	}

	// Return created user data (excluding the password)
	c.JSON(201, gin.H{"user": gin.H{
		"id":       user.ID,
		"username": user.Username,
	}})
}

// GetAllUsers retrieves all users
func GetAllUsers(c *gin.Context) {
	var users []models.User

	// Retrieve all users from the database
	if err := models.GetAllUsers(&users); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Return list of users
	c.IndentedJSON(http.StatusOK, &users)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	// Retrieve user by ID
	if err := models.GetUserByID(&user, id); err != nil {
		c.JSON(409, gin.H{"msg": "User Not Found"})
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Return user data
	c.IndentedJSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	// Delete user by ID
	if err := models.DeleteUser(&user, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Return a success message
	c.IndentedJSON(http.StatusOK, user)
}

// UpdateUser updates a user's details
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	// Find user by ID
	if err := models.FindFirstUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "User not in database!"})
		return
	}

	// Bind incoming JSON data to the user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": "invalid json input"})
		return
	}

	// Update user details
	if err := models.UpdateUser(&user, id); err != nil {
		c.AbortWithStatus(409) // Conflict error
		c.JSON(409, gin.H{"msg": "The request encountered a conflict"})
		return
	}

	// Return updated user data
	c.JSON(http.StatusOK, &user)
}

// Login function for user authentication
func Login(c *gin.Context) {
	var loginData models.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid JSON input"})
		return
	}

	// Check if the user exists
	var user models.User
	if err := models.GetUserByUsername(&user, loginData.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Invalid username or password",
		})
		return
	}

	// Check if the password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Invalid username or password",
		})
		return
	}

	// Generate JWT tokens
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET"),
		Issuer:            "AuthService",
		AccessExpiration:  720,    // Token expires in 720 minutes
		RefreshExpiration: 144000, // Refresh token expires in 1 day
	}

	accessToken, refreshToken, err := jwtWrapper.GenerateTokens(user.Username, user.ID)
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
	})
}

// RefreshToken handles the refresh token logic
func RefreshToken(c *gin.Context) {
	var tokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
		Username     string `json:"username" binding:"required"`
	}

	// Bind JSON data to the struct
	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid JSON input"})
		return
	}

	// Get user details from the database using the username
	var user models.User
	if err := models.GetUserByUsername(&user, tokenRequest.Username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	// Verify the refresh token validity
	valid, err := auth.VerifyRefreshTokenByUsername(tokenRequest.Username, tokenRequest.RefreshToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid or expired refresh token"})
		return
	}

	// Generate new access and refresh tokens
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET"),
		Issuer:            "AuthService",
		AccessExpiration:  30,     // Expire in 30 minutes
		RefreshExpiration: 144000, // Expire in 2400 hours or 100 days
	}
	accessToken, refreshToken, err := jwtWrapper.GenerateTokens(user.Username, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error generating tokens"})
		return
	}

	// Return the new tokens to the client
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
