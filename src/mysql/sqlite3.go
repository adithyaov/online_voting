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





