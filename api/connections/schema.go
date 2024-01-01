package connections

import (
	"dataspace/connections"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getSchemaRequest struct {
	ID int `json:"id"`
}

func handleSchema(c *gin.Context) {
	// Get the user ID from the context
	// username := c.GetString("username")
	var req getSchemaRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	schema, err := getSchema(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Schema",
		"schema":  schema,
	})
}

// getSchema returns the schema of the database
func getSchema(id int) (map[string]interface{}, error) {
	cnx := connections.ConnectionsManager.GetConnection(id) // Get the connection from the connection pool
	if cnx == nil {
		return nil, nil
	}
	return cnx.GetSchema()
}