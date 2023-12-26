package connections

import (
	"dataspace/db"
	"dataspace/db/types"

	"github.com/gin-gonic/gin"
)

func handleList(c *gin.Context) {
	username := c.GetString("username")
	connections := getConnections(username)

	connectionsRes := make(map[string]map[string]interface{})

	for _, connection := range connections {
		connectionsRes[connection.Host] = map[string]interface{}{
			"id":      connection.ID,
			"port":    connection.Port,
			"dbname":  connection.Dbname,
			"user":    connection.User,
			"SSLMode": connection.SSLMode,
		}
	}

	c.JSON(200, gin.H{"connections": connectionsRes})
}

// getConnections returns all connections for a given owner username
func getConnections(ownerUsername string) []types.Connection {
	var connections []types.Connection
	cnx := db.GetConnection()
	cnx.Select("id", "host", "port", "dbname", "user", "SSLMode").Find(&connections).Where("UserID = ?", cnx.Where("username = ?", ownerUsername).Select("id").Find(&types.User{}))
	return connections
}