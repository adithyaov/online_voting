package main

import (
	"fmt"
	"ballot"
	"sql"
)

func main() {
	fmt.Println("Hi testing! :-)")
	sql.CreateTables()
	sql.RunTransaction(sql.State{ballot.MakeBallot, []interface{}{"test", 5, 20, 22}})
}