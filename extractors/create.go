package extractors

import "database/sql"

type CRUDOperations interface {
	Insert(db *sql.DB, tableName string, columns []string, values []interface{}) (sql.Result, error)
	Select(db *sql.DB, tableName string, columns []string, where string, args ...interface{}) (*sql.Rows, error)
	Update(db *sql.DB, tableName string, columns []string, values []interface{}, where string, args ...interface{}) (sql.Result, error)
	Delete(db *sql.DB, tableName string, where string, args ...interface{}) (sql.Result, error)
	GetFields(db *sql.DB, tableName string) ([]string, error)
}
