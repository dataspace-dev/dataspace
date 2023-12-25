package connections

import "database/sql"

type connection struct {
	cnx *sql.DB
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
	p.cnx = cnx
	return nil
}