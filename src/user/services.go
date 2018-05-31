package user

import (
	c "common"
	"encoding/json"
	"net/http"
)

// DeleteAPI is the api to delete User
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

	json.NewEncoder(w).Encode(c.BasicResponse{Message: "Successfully deleted user",
		StatusCode: 200})

}
