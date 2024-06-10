package extractors

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// PostgreSQLCRUD реализует интерфейс CRUDInterface для PostgreSQL базы данных.
type PostgreSQLCRUD struct{}

// Create выполняет команду INSERT в PostgreSQL базе данных.
func (pc *PostgreSQLCRUD) Create(db *sql.DB, tableName string, columns []string, values []interface{}) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, joinColumns(columns), placeholders(len(values)))
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	return result, nil
}

// Read выполняет команду SELECT в PostgreSQL базе данных.
func (pc *PostgreSQLCRUD) Read(db *sql.DB, tableName string, columns []string, where string, args ...interface{}) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", joinColumns(columns), tableName, where)
	return db.Query(query, args...)
}

// Update выполняет команду UPDATE в PostgreSQL базе данных.
func (pc *PostgreSQLCRUD) Update(db *sql.DB, tableName string, columns []string, values []interface{}, where string, args ...interface{}) (sql.Result, error) {
	setClause := createSetClause(columns)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, setClause, where)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(append(values, args...)...)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	return result, nil
}

// Delete выполняет команду DELETE в PostgreSQL базе данных.
func (pc *PostgreSQLCRUD) Delete(db *sql.DB, tableName string, where string, args ...interface{}) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, where)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	return result, nil
}

// GetFields возвращает список полей таблицы в PostgreSQL базе данных.
func (pc *PostgreSQLCRUD) GetFields(db *sql.DB, tableName string) ([]string, error) {
	query := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name='%s'", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve columns: %v", err)
	}
	defer rows.Close()
	var fields []string
	for rows.Next() {
		var field string
		if err := rows.Scan(&field); err != nil {
			return nil, fmt.Errorf("could not scan column: %v", err)
		}
		fields = append(fields, field)
	}
	return fields, nil
}

// joinColumns соединяет имена колонок в строку через запятую.
func PostgresSQL_joinColumns(columns []string) string {
	return strings.Join(columns, ", ")
}

// placeholders создает строку с параметрами-заполнителями для SQL-запроса.
func PostgresSQL_placeholders(n int) string {
	ph := make([]string, n)
	for i := 0; i < n; i++ {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(ph, ", ")
}

// createSetClause создает строку SET для SQL-запроса UPDATE.
func PostgresSQL_createSetClause(columns []string) string {
	var setClauses []string
	for _, col := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, len(setClauses)+1))
	}
	return strings.Join(setClauses, ", ")
}
