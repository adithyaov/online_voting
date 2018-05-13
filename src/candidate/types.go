package candidate

import (
	"user"
	"ballot"
	"database/sql"
)

type Candidate struct {
	User *user.User         `json:"user"`
	Ballot *ballot.Ballot   `json:"ballot"`
	Details string 		    `json:"details"`
	Nominee1 sql.NullString `json:"nominee1"`
	Nominee2 sql.NullString `json:"nominee2"`
}



