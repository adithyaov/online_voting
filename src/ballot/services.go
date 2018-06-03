package ballot

import (
	"auth"
	c "common"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

// CreateAPI is an endpoint to create a ballot.
func CreateAPI(w http.ResponseWriter, r *http.Request, body *[]byte) {

	var data struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ballot, err := CreateBallot(data.Code, data.Name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(ballot)

}

// FindAPI is an endpoint to get ballot information.
func FindAPI(w http.ResponseWriter, r *http.Request, body *[]byte) {

	var data struct {
		Code string `json:"code"`
	}

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	b, err := OpenBallot(data.Code)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(*b)

}

// DeleteAPI creates an endpoint to delete ballot
func DeleteAPI(w http.ResponseWriter, r *http.Request, body *[]byte) {

	var data struct {
		Code string `json:"code"`
	}

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = DeleteBallot(data.Code)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(c.BasicResponse{
		Message:    "Successfully deleted the ballot",
		StatusCode: 200})

}

// BlindVoteAPI provides an endpoint to blind the vote
func BlindVoteAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

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

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	bias := strconv.FormatFloat((rand.Float64()*100000)+rand.Float64(), 'f', 6, 64)
	vote := Vote{ballot.Code, data.CandidateEmail, bias}
	hashed, err := vote.Hash()

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	blinded, unblinder, err := ballot.BlindVote(vote)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	response := Res{c.ConvertBSToIS(blinded), c.ConvertBSToIS(unblinder), c.ConvertBSToIS(hashed), bias}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SignBytesAPI provides an endpoint so sign with a specific ballot
func SignBytesAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

	type Req struct {
		Blinded    []int  `json:"blinded"`
		VoterEmail string `json:"voter_email"`
	}

	type Res struct {
		Signed []int `json:"signed"`
	}

	var data Req

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	signed, err := ballot.SignBlindHash(c.ConvertISToBS(data.Blinded))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Before responding Note the token, save the token. Auth Field required.
	token := r.Header["token"][0]
	gt, err := auth.ParseToken(token)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if !(gt.Email == data.VoterEmail || gt.RoleCode == "A") {
		http.Error(w, "Not permitted.", 400)
		return
	}

	response := Res{c.ConvertBSToIS(signed)}

	w.Header().Set("Content-Type", "application/json")

	if err := ballot.AddVoter(data.VoterEmail); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// UnblindSignAPI provides an endpoint to unblind the sign
func UnblindSignAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

	type Req struct {
		Signed    []int `json:"signed"`
		Unblinder []int `json:"unblinder"`
	}

	type Res struct {
		Unblinded []int `json:"unblinded"`
	}

	var data Req

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	unblinded := ballot.UnblindSignedHash(c.ConvertISToBS(data.Signed), c.ConvertISToBS(data.Unblinder))

	response := Res{c.ConvertBSToIS(unblinded)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// VerifySignAPI provides a way to check if the sign given is proper
func VerifySignAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

	type Req struct {
		Hashed    []int `json:"vote_hash"`
		Unblinded []int `json:"unblinded"`
	}

	type Res struct {
		Err string `json:"error"`
	}

	var data Req

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = ballot.VerifySign(c.ConvertISToBS(data.Hashed), c.ConvertISToBS(data.Unblinded))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Res{""})
}

// FindBallotsAPI returns all ballots depending on the user
func FindBallotsAPI(w http.ResponseWriter, r *http.Request, openBallots map[string]*Ballot) {

	token := r.Header["token"][0]
	gt, err := auth.ParseToken(token)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(GetBallots(gt.Email, openBallots))

}
