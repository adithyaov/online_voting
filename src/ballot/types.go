package ballot

import (
	"math/big"
)


type Ballot struct {
	Code string
	Name string
	N big.Int
	D big.Int
	E int
	Flag bool
}


type Vote struct {
	BallotCode string `json:"ballot_code"`
	CandidateEmail string `json:"candidate_email"`
	Bias string `json:"bias"`
}






