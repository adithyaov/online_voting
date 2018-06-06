package candidate

import (
	"auth"
	"ballot"
	"encoding/json"
	"user"
)

// CreateAPI provides the endpoint for createing a candidate
func CreateAPI(s auth.Service) {
	type Req struct {
		BallotCode string `json:"ballot_code"`
		UserEmail  string `json:"user_email"`
		Details    string `json:"details"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	ballot, err := ballot.OpenBallot(data.BallotCode)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	user, err := user.GetUser(data.UserEmail)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	candidate := Candidate{Ballot: ballot, User: user, Details: data.Details}
	err = candidate.Create()
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(candidate, 200)
}

// AddNomineeAPI adds a nominee
func AddNomineeAPI(s auth.Service) {
	type Req struct {
		BallotCode   string `json:"ballot_code"`
		UserEmail    string `json:"user_email"`
		NomineeEmail string `json:"nominee_email"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	candidate, err := GetCandidate(data.BallotCode, data.UserEmail)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	if !candidate.Nominee1.Valid {
		candidate.Nominee1.Valid = true
		candidate.Nominee1.String = data.NomineeEmail
	} else if !candidate.Nominee1.Valid {
		candidate.Nominee2.Valid = true
		candidate.Nominee2.String = data.NomineeEmail
	} else {
		json.NewEncoder(s.Writer).Encode(candidate)
		return
	}

	err = candidate.UpdateNominees()
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(candidate, 200)
}

// UpdateDetailsAPI gives an endpoint to update details
func UpdateDetailsAPI(s auth.Service) {
	type Req struct {
		BallotCode string `json:"ballot_code"`
		UserEmail  string `json:"user_email"`
		Details    string `json:"details"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	candidate, err := GetCandidate(data.BallotCode, data.UserEmail)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	candidate.Details = data.Details

	err = candidate.UpdateDetails()
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(candidate, 200)
}

// DeleteAPI gives an endpoint to delete candidate
func DeleteAPI(s auth.Service) {
	type Req struct {
		BallotCode string `json:"ballot_code"`
		UserEmail  string `json:"user_email"`
	}

	var data Req

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	DeleteCandidate(data.BallotCode, data.UserEmail)
	s.Tell("Successfully deleted candidate", 200)
}
