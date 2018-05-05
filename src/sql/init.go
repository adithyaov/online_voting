package sql

import (
	"log"
	"ballot"
)

func CreateTables() {
	RunRawString(ballot.BallotTable)
	log.Printf("Created tables")
}