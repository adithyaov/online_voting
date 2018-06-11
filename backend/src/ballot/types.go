package ballot

import (
	"auth"
	"math/big"
	"net/http"
)

// Ballot describes a ballot
type Ballot struct {
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	N               big.Int `json:"n"`
	D               big.Int `json:"-"`
	E               int     `json:"e"`
	RegexpVoter     string  `json:"regex_voter"`
	RegexpCandidate string  `json:"regex_candidate"`
	Phase           string  `json:"phase"`
}

// Vote defines a vote structure
type Vote struct {
	BallotCode     string `json:"ballot_code"`
	CandidateEmail string `json:"candidate_email"`
	Bias           string `json:"bias"`
}

// BodyService is the expected type of service which needs a ballot and a body
type BodyService func(http.ResponseWriter, *http.Request, *Ballot, *[]byte)

// OpenBallotService is the ecpected type of service which needs a ballot
type OpenBallotService func(http.ResponseWriter, *http.Request, map[string]*Ballot)

// Service is the expected input type for all services
type Service struct {
	auth.Service
	Ballot      *Ballot
	OpenBallots map[string]*Ballot
}
