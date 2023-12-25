package auth

import (
	"dataspace/db"
	"dataspace/db/models"

	"github.com/gin-gonic/gin"
)

type signupRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signupHandler(c *gin.Context) {

	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.Username == "" || req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	conflict, err := createUser(req)
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

func createUser(req signupRequest) (conflict bool, err error) {
	cnx := db.GetConnection()

	var user models.User = models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	exists := cnx.Where("username = ?", user.Username).Or("email = ?", user.Email).First(&user).Error == nil
	if exists {
		return true, nil
	}

	err = cnx.Create(&user).Error
	if err != nil {
		return false, err
	}

	return false, nil
}
