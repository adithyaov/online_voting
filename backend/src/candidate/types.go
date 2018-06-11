package candidate

import (
	"ballot"
	"database/sql"
	"user"
)

// Candidate describes a basic candidate
type Candidate struct {
	User     *user.User     `json:"user"`
	Ballot   *ballot.Ballot `json:"ballot"`
	Details  string         `json:"details"`
	Nominee1 sql.NullString `json:"nominee1"`
	Nominee2 sql.NullString `json:"nominee2"`
}

// PartialCandidate describes a basic candidate without the entire ballot
type PartialCandidate struct {
	User       *user.User     `json:"user"`
	BallotCode string         `json:"ballot_code"`
	Details    string         `json:"details"`
	Nominee1   sql.NullString `json:"nominee1"`
	Nominee2   sql.NullString `json:"nominee2"`
}
