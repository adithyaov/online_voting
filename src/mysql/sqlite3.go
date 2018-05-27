package mysql

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Side effect registers sqlite3
)

// RunRawString runs a raw SQL string.
func RunRawString(rawString string) error {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(rawString)
	if err != nil {
		return err
	}

	return nil
}

// OpenDB opens a pointer to DB, should be used carefully.
func OpenDB() (*sql.DB, error) {
	return sql.Open("sqlite3", SqliteStore)
}

// Exec executes a query.
func Exec(query string, args []interface{}) (*(sql.Result), error) {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// QueryOne queries for 1 result and reads them to the respective.
func QueryOne(query string, args []interface{}, scanTo []interface{}) error {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		return err
	}
	err = db.QueryRow(query, args...).Scan(scanTo...)
	db.Close()
	return err
}
