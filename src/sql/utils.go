package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateTable(storePath string, schema string) {
	db, err := sql.Open("sqlite3", storePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(schema)
	if err != nil {
		log.Printf("%q: %s\n", err, schema)
		return
	}
	log.Printf("Table created %s", schema)
}


func RunTransaction(storePath string, stmt string, stmtParams []interface{}) {
	db, err := sql.Open("sqlite3", storePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	prepareStmt, err := tx.Prepare(stmt)
	defer prepareStmt.Close()

	_, err = prepareStmt.Exec(stmtParams...)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}



