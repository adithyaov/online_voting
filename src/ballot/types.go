package ballot

import (
	"math/big"
)


type Ballot struct {
	Code string 		  `json:"code"`
	Name string 		  `json:"name"`
	N big.Int 			  `json:"n"`
	D big.Int 			  `json:"-"`
	E int 				  `json:"e"`
	RegexVoter string 	  `json:"regex_voter"`
	RegexCandidate string `json:"regex_candidate"`
	Phase string 	      `json:"phase"`
}


type Vote struct {
	BallotCode string     `json:"ballot_code"`
	CandidateEmail string `json:"candidate_email"`
	Bias string  		  `json:"bias"`
}



