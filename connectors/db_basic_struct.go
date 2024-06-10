package connectors

import "database/sql"

type DBConnector interface {
	Connect(username, password, hostname, port, dbname string) (*sql.DB, error)
}
