package candidate

import (
	"encoding/json"
	"net/http"
)

// CreateCandidateAPI provides the endpoint for createing a candidate
func CreateCandidateAPI(w http.ResponseWriter, r *http.Request, body *[]byte) {
	type Req struct {
		BallotCode []string `json:"ballot_code"`
		UserEmail  []string `json:"user_email"`
	}

	type Res struct {
		Error error `json:"error"`
	}

	var data Req

	err := json.Unmarshal(*body, &data)

}
