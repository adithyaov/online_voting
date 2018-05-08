package candidate

import (
	"user"
	"ballot"
)

type Candidate struct {
	User *user.User       `json:"user"`
	Ballot *ballot.Ballot `json:"ballot"`
	Details string 		  `json:"details"`
}



