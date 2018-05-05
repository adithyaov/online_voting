package ballot

import (
	"net/http"
	"encoding/json"
	//"mysql"
)


func CreateAPI(w http.ResponseWriter, r *http.Request) {

	var data struct {
		code string
		name string
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ballot, err := CreateBallot(data.code, data.name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(ballot)

}

// func FindAPI(w http.ResponseWriter, r *http.Request) {

// 	var data struct {
// 		code string
// 	}

// 	if r.Body == nil {
// 		http.Error(w, "Please send a request body", 400)
// 		return
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}

// 	row, err := mysql.RunTransaction(mysql.State{GetBallot, []interface{}{data.code}})
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}

// 	var ballot Ballot
// 	row.Scan(&ballot.code, &ballot.name, &ballot.n, &ballot.d, &ballot.e, &ballot.flag)

// 	json.NewEncoder(w).Encode(ballot)

// }
