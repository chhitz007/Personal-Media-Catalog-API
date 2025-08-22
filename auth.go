package controllers

import (
	"bookapi/config"              // DB config
	"bookapi/models"              // User model
	"net/http"                    // HTTP status codes
	"os"                          // Access environment variables
	"time"                        // Token expiration

	"github.com/gin-gonic/gin"    // Gin web framework
	"github.com/golang-jwt/jwt/v5" // JWT generation
	"golang.org/x/crypto/bcrypt"  // Password hashing
)

// Register handles user signup
func Register(c *gin.Context) {
	var user models.User // Declare a user variable to hold incoming JSON

	// Parse JSON input into user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return error if input is invalid
		return
	}

	// Hash the user's password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"}) // Hashing error
		return
	}
	user.Password = string(hashedPassword) // Replace plain password with hashed version

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"}) // Handle duplicate username
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user authentication and returns JWT
func Login(c *gin.Context) {
	var input models.User  // Input struct for incoming credentials
	var dbUser models.User // dbUser will hold user fetched from DB

	// Parse input JSON into input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Invalid input
		return
	}

	// Look for a user with matching username in DB
	if err := config.DB.Where("username = ?", input.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // User not found
		return
	}

	// Check if provided password matches hashed password in DB
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // Wrong password
		return
	}

	// Create a new JWT token with user ID and expiry
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dbUser.ID,                          // Store user ID in token
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Set expiration to 24 hours
	})

	// Load JWT secret key from environment variable
	secret := os.Getenv("JWT_SECRET")

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"}) // Token signing failed
		return
	}

	// Return the token to the client 
	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
	})
}

