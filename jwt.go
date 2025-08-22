package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware validates incoming JWT tokens and extracts the user_id claim
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the Authorization header from the request
		authHeader := c.GetHeader("Authorization")
		fmt.Println("🔍 Incoming Authorization header:", authHeader)

		// Check if the header exists and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("❌ Missing or malformed Authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token missing or malformed",
			})
			return
		}

		// Remove "Bearer " prefix and trim any extra spaces
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		fmt.Println("🔑 Extracted Token String:", tokenString)

		// Get the JWT secret from environment variable
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			fmt.Println("⚠️ WARNING: JWT_SECRET environment variable is not set!")
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC (security check)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("❌ Unexpected signing method:", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			// If parsing fails, reject the request
			fmt.Println("❌ Token parsing error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check if token is valid (not expired or tampered)
		if !token.Valid {
			fmt.Println("❌ Token is invalid or expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// ✅ Extract claims (payload) from token and set user_id in Gin context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println("✅ Token Claims:", claims)

			// user_id is stored as float64 in MapClaims, convert to uint
			if userID, ok := claims["user_id"].(float64); ok {
				fmt.Printf("👉 Extracted user_id: %v\n", uint(userID))
				c.Set("user_id", uint(userID)) // set user_id for use in controllers
			} else {
				fmt.Println("⚠️ user_id claim missing or wrong type")
			}
		} else {
			fmt.Println("⚠️ Failed to cast token claims to MapClaims")
		}

		// Continue to the next handler
		c.Next()
	}
}


