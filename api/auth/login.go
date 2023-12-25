package auth

import (
	"dataspace/db"
	"dataspace/db/models"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	// Validate the request
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	user, err := getUser(req.Username) // Get the user
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if user == nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	} else if !checkPassword(req.Password, user.Password) {
		c.JSON(401, gin.H{
			"message": "Incorrect password",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Login successful",
			"bearer token": generateToken(user),
		})
	}
}

// checkPassword checks if the password matches the hash
// It returns true if the password matches the hash, false otherwise
func checkPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// getUser gets a user from the database
// Returns nil if the user is not found
func getUser(username string) (*models.User, error) {
	var user models.User
	result := db.GetConnection().Where("username = ?", username).Find(&user)
	if result.Error != nil {
		fmt.Println("Error getting user:", result.Error)
		return nil, fmt.Errorf("error getting user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

// generateToken generates a JWT token for the user
func generateToken(user *models.User) (string) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET environment variable not set")	
		os.Exit(1)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}
	return tokenString
}