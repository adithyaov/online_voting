package user

import (
	"auth"
	"encoding/json"
)

// DeleteAPI is the api to delete User
func DeleteAPI(s auth.Service) {

	var data struct {
		Email string `json:"email"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	err = DeleteUser(data.Email)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Tell("Successfully deleted user", 200)

}
