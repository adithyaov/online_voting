package main

import (
	"fmt"
	"ballot"
	"sql"
)

func main() {
	fmt.Println("Hi testing! :-)")
	sql.CreateTable(ballot.SqliteStore, ballot.BallotTable)
	sql.RunTransaction(ballot.SqliteStore, ballot.MakeBallot, []interface{}{"test", 5, 20, 22})
}