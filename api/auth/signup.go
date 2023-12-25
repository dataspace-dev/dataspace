package auth

import (
	"dataspace/db"
	"dataspace/db/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signupHandler(c *gin.Context) {
	// Validate the request
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.Username == "" || req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	encryptPassword(&req.Password) // Encrypt the password

	conflict, err := createUser(req) // Create the user
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if conflict {
		c.JSON(409, gin.H{
			"message": "Username or email already exists",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "User created",
		})
		return
	}
}

// encryptPassword encrypts the password using bcrypt
// It modifies the password string in place
func encryptPassword(password *string) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	*password = string(hashed)
}

// createUser creates a new user in the database
// It returns a boolean indicating whether there was a conflict and an error
func createUser(req signupRequest) (conflict bool, err error) {
	cnx := db.GetConnection()

	var user models.User = models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	existingUser := cnx.Find(&user, "username = ? OR email = ?", req.Username, req.Email)
	if existingUser.RowsAffected > 0 {
		return true, nil
	}	

	err = cnx.Create(&user).Error
	if err != nil {
		return false, err
	}

	return false, nil
}
