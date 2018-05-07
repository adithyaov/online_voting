package ballot

import (
	"net/http"
	"encoding/json"
	"math/rand"
	"mysql"
	"strconv"
	c "common"
)


func CreateAPI(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Code string
		Name string
	}

	err := json.NewDecoder(r.Body).Decode(&data)
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

func FindAPI(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Code string
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	rows, err := mysql.RunQuery(mysql.State{GetBallot, []interface{}{data.Code}})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var ballot Ballot

	rows.Next()
	var n, d string
	rows.Scan(&ballot.Code, &ballot.Name, &n, &d, &ballot.E, &ballot.Flag)
	ballot.N.SetString(n, 10)
	ballot.D.SetString(d, 10)

	json.NewEncoder(w).Encode(ballot)

}


func DeleteAPI(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Code string
	}

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = mysql.RunTransaction(mysql.State{DeleteBallot, []interface{}{data.Code}})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(BasicResponse{"Successfully deleted the ballot", 200})

}



func BlindVoteAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

	type Req struct {
		CandidateEmail string `json:"candidate_email"`
	}

	type Res struct {
		Blinded []int  	`json:"blinded"`
		Unblinder []int `json:"unblinder"`
		VoteHash []int  `json:"vote_hash"`
		Bias string 	`json:"bias"`
	}

	var data Req

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	bias := strconv.FormatFloat((rand.Float64() * 100000) + rand.Float64(), 'f', 6, 64)
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



func SignBytesAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

	type Req struct {
		Blinded []int `json:"blinded"`
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


	response := Res{c.ConvertBSToIS(signed)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}





func UnblindSignAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot,  body *[]byte) {

	type Req struct {
		Signed []int    `json:"signed"`
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



func VerifySignAPI(w http.ResponseWriter, r *http.Request, ballot *Ballot, body *[]byte) {

	type Req struct {
		Hashed []int      `json:"vote_hash"`
		Unblinded []int   `json:"unblinded"`
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

	response := Res{""}
	if err != nil {
		response = Res{err.Error()}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}













