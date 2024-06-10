package connectors

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnector struct{}

func (c *MySQLConnector) Connect(username, password, hostname, port, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("%d: could not open db: %v", 1, err)
	}

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%d: could not connect to db: %v", 2, err)
	}

	fmt.Println("Connected to MySQL successfully!")
	return db, nil
}
