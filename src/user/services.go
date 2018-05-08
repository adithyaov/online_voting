package user

import (
	"net/http"
	"encoding/json"
	c "common"
)

func DeleteAPI(w http.ResponseWriter, r *http.Request, body *[]byte) {

	var data struct {
		Email string `json:"email"`
	}

	err := json.Unmarshal(*body, &data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = DeleteUser(data.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}


	json.NewEncoder(w).Encode(c.BasicResponse{"Successfully deleted user", 200})

}







