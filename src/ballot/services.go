package ballot

import (
	"net/http"
	"encoding/json"
	"math/rand"
	"mysql"
	"strconv"
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


func BlindVoteAPI(openBallots []*Ballot) http.HandlerFunc {


	return func (w http.ResponseWriter, r *http.Request) {

		var data struct {
			BallotCode string     `json:"ballot_code"`
			CandidateEmail string `json:"candidate_email"`
		}

		var err error

		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}


		var ballot *Ballot
		ballot = nil
		for _, b := range openBallots {
			if b.Code == data.BallotCode {
				ballot = b
				break
			}
		}
		if ballot == nil {
			http.Error(w, "Ballot not found", 400)
			return
		}



		bias := strconv.FormatFloat((rand.Float64() * 10000) + rand.Float64(), 'f', 6, 64)
		vote := Vote{data.BallotCode, data.CandidateEmail, bias}
		

		blinded, unblinder, err := ballot.BlindVote(vote)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		var intBlinded []int
		var intUnblinder []int

		for _, b := range blinded {
			intBlinded = append(intBlinded, int(b))
		}

		for _, b := range unblinder {
			intUnblinder = append(intUnblinder, int(b))
		}

		response := struct {
			Blinded []int  	`json:"blinded"`
			Unblinder []int `json:"unblinder"`
			Bias string 	`json:"bias"`
		}{intBlinded, intUnblinder, bias}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}