package mysql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func RunRawString(rawString string) {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(rawString)
	if err != nil {
		log.Printf("%q: %s\n", err, rawString)
		return
	}
}


func RunTransaction(state State) {
	db, err := sql.Open("sqlite3", SqliteStore)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	prepareStmt, err := tx.Prepare(state.Stmt)
	defer prepareStmt.Close()

	_, err = prepareStmt.Exec(state.Params...)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}


// func RunTransactionsCS(states []State) {
// 	db, err := sql.Open("sqlite3", SqliteStore)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	for _, state := range states {
// 		tx, err := db.Begin()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		prepareStmt, err := tx.Prepare(state.Stmt)
// 		defer prepareStmt.Close()

// 		_, err = prepareStmt.Exec(state.Params...)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tx.Commit()
// 		log.Printf("%s: %v\n", state.Stmt, state.Params)
// 	}
// }

// func RunTransactionsCA(states []State) {
// 	db, err := sql.Open("sqlite3", SqliteStore)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
	
// 	tx, err := db.Begin()
// 	for _, state := range states {
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		prepareStmt, err := tx.Prepare(state.Stmt)
// 		defer prepareStmt.Close()

// 		_, err = prepareStmt.Exec(state.Params...)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("%s: %v\n", state.Stmt, state.Params)
// 	}
// 	tx.Commit()
// }





