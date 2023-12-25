package connections

import "dataspace/db/models"

type connectionPool struct {
	pool map[int]*connection
}

// GetConnection retrieves a connection from the connection pool based on the given ID.
// It returns the connection if found, otherwise it returns nil.
func (p *connectionPool) GetConnection(id int) *connection {
	return p.pool[id]
}

// AddConnection adds a connection to the connection pool.
func (p *connectionPool) AddConnection(id int, dbData *models.Conection) error {
	dsn := "host=" + dbData.Host + " port=" + dbData.Port + " dbname=" + dbData.Dbname + " user=" + dbData.User + " password=" + dbData.Pass + " sslmode=" + dbData.SSLMode
	var cnx connection
	err := cnx.Connect(dsn)
	if err != nil {
		return err
	}
	p.pool[id] = &cnx
	return nil
}
