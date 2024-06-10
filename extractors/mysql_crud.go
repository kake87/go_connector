package extractors

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLCRUD реализует интерфейс CRUDInterface для MySQL базы данных.
type MySQLCRUD struct{}

// Create выполняет команду INSERT в MySQL базе данных.
func (mc *MySQLCRUD) Create(db *sql.DB, tableName string, columns []string, values []interface{}) (sql.Result, error) {
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

// Read выполняет команду SELECT в MySQL базе данных.
func (mc *MySQLCRUD) Read(db *sql.DB, tableName string, columns []string, where string, args ...interface{}) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", joinColumns(columns), tableName, where)
	return db.Query(query, args...)
}

// Update выполняет команду UPDATE в MySQL базе данных.
func (mc *MySQLCRUD) Update(db *sql.DB, tableName string, columns []string, values []interface{}, where string, args ...interface{}) (sql.Result, error) {
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

// Delete выполняет команду DELETE в MySQL базе данных.
func (mc *MySQLCRUD) Delete(db *sql.DB, tableName string, where string, args ...interface{}) (sql.Result, error) {
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

// GetFields возвращает список полей таблицы в MySQL базе данных.
func (mc *MySQLCRUD) GetFields(db *sql.DB, tableName string) ([]string, error) {
	query := fmt.Sprintf("SHOW COLUMNS FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve columns: %v", err)
	}
	defer rows.Close()
	var fields []string
	for rows.Next() {
		var field, dataType, null, key, defaultValue, extra string
		if err := rows.Scan(&field, &dataType, &null, &key, &defaultValue, &extra); err != nil {
			return nil, fmt.Errorf("could not scan column: %v", err)
		}
		fields = append(fields, field)
	}
	return fields, nil
}

// joinColumns соединяет имена колонок в строку через запятую.
func joinColumns(columns []string) string {
	return strings.Join(columns, ", ")
}

// placeholders создает строку с параметрами-заполнителями для SQL-запроса.
func placeholders(n int) string {
	return strings.TrimSuffix(strings.Repeat("?, ", n), ", ")
}

// createSetClause создает строку SET для SQL-запроса UPDATE.
func createSetClause(columns []string) string {
	var setClauses []string
	for _, col := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
	}
	return strings.Join(setClauses, ", ")
}
