package connections

import (
	"dataspace/db"
	"dataspace/db/types"
	"fmt"
	"strings"
)

// The global connection manager
var ConnectionsManager = manager{
	pool: make(map[int]*connection),
}

// This is the type of the connection manager
type manager struct {
	// pool is a map that stores the connections
	pool map[int]*connection
}

// GetConnection retrieves a connection from the connection pool based on the given ID.
// It returns the connection if found, otherwise it returns nil.
func (m *manager) GetConnection(id int) *connection {
	if _, ok := m.pool[id]; !ok { // Check if the connection does not exist in the pool
		cnx := db.GetConnection()
		var dbData types.Connection
		if err := cnx.Find(&dbData, "id = ?", id).Error; err != nil { // Get the connection data from the database
			return nil
		}
		fmt.Printf("Pool size: %d\n", len(m.pool))
		if err := m.AddConnection(id, &dbData); err != nil { // Add the connection to the connection pool
			return nil
		}
		fmt.Printf("Added connection to the pool (id: %d)\n", id)
		fmt.Printf("Pool size: %d\n", len(m.pool))
	}
	return m.pool[id]
}

// AddConnection adds a connection to the connection pool.
func (m *manager) AddConnection(id int, dbData *types.Connection) error {
	dsn := BuildDsn(dbData)
	var cnx connection
	err := cnx.Connect(dsn)
	if err != nil {
		return fmt.Errorf("there was an error connecting to the database (id: %d): %w", id, err)
	}
	cnx.SetOwner(dbData.UserID)
	m.pool[id] = &cnx
	return nil
}

// TestConnectionTemporarily tests a connection to the database given de connection data without adding it to the connection pool.
// It returns an error if the connection fails.
func (m *manager) TestConnectionTemporarily(dbData *types.Connection) error {
	dsn := BuildDsn(dbData)
	var cnx connection
	err := cnx.Connect(dsn)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			return fmt.Errorf("the connection was refused, please check the connection data")
		} else if strings.Contains(err.Error(), "no pg_hba.conf entry for host") {
			return fmt.Errorf("the host is not allowed to connect to the database, please check the connection data")
		} else if strings.Contains(err.Error(), "password authentication failed") {
			return fmt.Errorf("the password is incorrect, please check the connection data")
		} else {
			return fmt.Errorf("there was an error connecting to the database: %w", err)
		}
	}
	err = cnx.Ping()
	if err != nil {
		return fmt.Errorf("there was an error pinging the database: %w", err)
	}
	return nil
}
