package connections

import "database/sql"

type connection struct {
	// The raw connection to the database
	cnx *sql.DB

	// We store the DSN in the connection struct so that we can use it to reconnect
	dsn string
}

// Ping sends a ping request to the database server to check the connection status.
// It returns an error if the ping fails.
func (p *connection) Ping() error {
	return p.cnx.Ping()
}

func (p *connection) Connect(dsn string) error {
	cnx, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	p = &connection{
		cnx: cnx,
		dsn: dsn,
	}
	return nil
}