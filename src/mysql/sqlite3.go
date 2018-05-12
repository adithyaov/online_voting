package mysql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

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


func RunTransaction(state State) error {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	prepareStmt, err := tx.Prepare(state.Stmt)
	defer prepareStmt.Close()

	_, err = prepareStmt.Exec(state.Params...)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func RunQuery(state State) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		return nil, err
	}
	defer db.Close()


	return db.Query(state.Stmt, state.Params...)

}


func OpenDB() (*sql.DB, error) {
	return sql.Open("sqlite3", SqliteStore)
}

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

func QueryOne(query string, args []interface{}, scanTo []interface{}) error {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		return err
	}
	err = db.QueryRow(query, args...).Scan(scanTo...)
	db.Close()
	return err
}



