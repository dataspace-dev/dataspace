package connections

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type connection struct {
	// The raw connection to the database
	cnx *sql.DB

	// We store the DSN in the connection struct so that we can use it to reconnect
	dsn string
}

// Ping sends a ping request to the database server to check the connection status.
// It returns an error if the ping fails.
func (p *connection) Ping() error {
	if p.cnx == nil {
		return sql.ErrConnDone
	}
	return p.cnx.Ping()
}

// Connect connects to the database using the given DSN.
// It returns an error if the connection fails.
func (p *connection) Connect(dsn string) error {
	cnx, err := sql.Open("postgres", dsn)
	if err != nil { // We try to connect to the database
		return err
	}
	if err := cnx.Ping(); err != nil { // We ping it to check if the connection is valid
		return err
	}
	p.cnx = cnx // Update the cnx field of the receiver
	p.dsn = dsn // Update the dsn field of the receiver
	return nil
}
