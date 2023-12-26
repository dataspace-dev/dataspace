package connections

import (
	"dataspace/connections"
	"dataspace/db"
	"dataspace/db/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
}

func handleCreate(c *gin.Context) {
	// Validate the request
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Host == "" || req.Port == "" || req.Dbname == "" || req.User == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var dbData = &types.Connection{
		Host:    req.Host,
		Port:    req.Port,
		Dbname:  req.Dbname,
		User:    req.User,
		Pass:    req.Password,
		SSLMode: req.SSLMode,
	}

	if dbData.SSLMode == "" {
		dbData.SSLMode = "disable"
	}

	if err := testConnection(dbData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid connection data",
			"error":   err.Error(),
		})
		return
	}

	id, err := saveConnection(dbData, c.GetString("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error saving connection",
			"error":   err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection saved",
			"id": id,
		})
	}
}

// testConnection tests a connection to the database given de connection data.
func testConnection(data *types.Connection) error {
	return connections.ConnectionsManager.TestConnectionTemporarily(data)
}

// saveConnection saves a connection to the database given de connection data.
func saveConnection(data *types.Connection, ownerUsername string) (id int, err error) {
	cnx := db.GetConnection()

	// we need to get the owner's ID
	var user types.User
	cnx.First(&user, "username = ?", ownerUsername)

	var connection = types.Connection{
		Host:    data.Host,
		Port:    data.Port,
		Dbname:  data.Dbname,
		User:    data.User,
		Pass:    data.Pass,
		SSLMode: data.SSLMode,
		UserID:  user.ID,
	}

	insert := cnx.Create(&connection)
	if insert.Error != nil {
		return 0, insert.Error
	}
	return int(connection.ID), nil
}
