package ballot

import (
	"auth"
	c "common"
	"encoding/json"
	"math/rand"
	"strconv"
)

// CreateAPI is an endpoint to create a ballot.
func CreateAPI(s auth.Service) {

	var data struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	ballot, err := CreateBallot(data.Code, data.Name)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(ballot, 200)

}

// FindAPI is an endpoint to get ballot information.
func FindAPI(s auth.Service) {

	var data struct {
		Code string `json:"code"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	b, err := OpenBallot(data.Code)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(b, 200)

}

// DeleteAPI creates an endpoint to delete ballot
func DeleteAPI(s auth.Service) {

	var data struct {
		Code string `json:"code"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	err = DeleteBallot(data.Code)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Tell("Successfully deleted the ballot", 200)

}

// BlindVoteAPI provides an endpoint to blind the vote
func BlindVoteAPI(s Service) {

	type Req struct {
		CandidateEmail string `json:"candidate_email"`
	}

	type Res struct {
		Blinded   []int  `json:"blinded"`
		Unblinder []int  `json:"unblinder"`
		VoteHash  []int  `json:"vote_hash"`
		Bias      string `json:"bias"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	if ok, err := s.Ballot.IsCandidate(data.CandidateEmail); err == nil {
		if ok == false {
			s.Tell("Candidate not a part of Ballot", 400)
			return
		}
	} else {
		s.Tell(err.Error(), 400)
		return
	}

	bias := strconv.FormatFloat((rand.Float64()*100000)+rand.Float64(), 'f', 6, 64)
	vote := Vote{s.Ballot.Code, data.CandidateEmail, bias}
	hashed, err := vote.Hash()

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	blinded, unblinder, err := s.Ballot.BlindVote(vote)

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	response := Res{c.ConvertBSToIS(blinded),
		c.ConvertBSToIS(unblinder),
		c.ConvertBSToIS(hashed), bias}

	s.Encode(response, 200)
}

// SignBytesAPI provides an endpoint so sign with a specific ballot
func SignBytesAPI(s Service) {

	type Req struct {
		Blinded    []int  `json:"blinded"`
		VoterEmail string `json:"voter_email"`
	}

	type Res struct {
		Signed []int `json:"signed"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	signed, err := s.Ballot.SignBlindHash(c.ConvertISToBS(data.Blinded))

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	// Before responding Note the token, save the token. Auth Field required.

	if !(s.Token.Email == data.VoterEmail || c.IsIn(s.Token.RoleCode, "A")) {
		s.Tell("Not permitted", 400)
		return
	}

	response := Res{c.ConvertBSToIS(signed)}

	if err := s.Ballot.AddVoter(data.VoterEmail); err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(response, 200)
}

// UnblindSignAPI provides an endpoint to unblind the sign
func UnblindSignAPI(s Service) {

	type Req struct {
		Signed    []int `json:"signed"`
		Unblinder []int `json:"unblinder"`
	}

	type Res struct {
		Unblinded []int `json:"unblinded"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	unblinded := s.Ballot.UnblindSignedHash(c.ConvertISToBS(data.Signed),
		c.ConvertISToBS(data.Unblinder))

	response := Res{c.ConvertBSToIS(unblinded)}

	s.Encode(response, 200)
}

// VerifySignAPI provides a way to check if the sign given is proper
func VerifySignAPI(s Service) {

	type Req struct {
		Hashed    []int `json:"vote_hash"`
		Unblinded []int `json:"unblinded"`
	}

	type Res struct {
		Err string `json:"error"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	err = s.Ballot.VerifySign(c.ConvertISToBS(data.Hashed),
		c.ConvertISToBS(data.Unblinded))

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(Res{""}, 200)
}

// FindBallotsAPI returns all ballots depending on the user
func FindBallotsAPI(s Service) {

	type Req struct {
		Email string `json:"email"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(GetBallots(data.Email, s.OpenBallots), 200)

}
