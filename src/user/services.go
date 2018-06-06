package user

import (
	"auth"
	c "common"
	"encoding/json"
)

// AuthUserAPI is the api to create User
func AuthUserAPI(s auth.Service) {

	type Res struct {
		Token  string `json:"token"`
		Status string `json:"status"`
	}

	token := s.Request.Header["GoogleToken"]
	if len(token) == 0 {
		s.Tell("No GoogleToken found in Header", 400)
		return
	}

	jwtToken, err := auth.GenerateToken(token[0])
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	gt, err := auth.ParseToken(jwtToken)

	user := User{}
	user.FromToken(gt)

	ok, err := user.CheckIfExists()
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	if ok {
		s.Encode(Res{jwtToken, "sustained"}, 200)
		return
	}

	err = user.Create()
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(Res{jwtToken, "created"}, 200)

}

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

	if !(data.Email == s.Token.Email || c.IsIn(s.Token.RoleCode, "A")) {
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

// UpdatePersonalAPI is the api to update name and/or picture
func UpdatePersonalAPI(s auth.Service) {

	var data struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	// Ownership or Admin
	if !(data.Email == s.Token.Email || c.IsIn(s.Token.RoleCode, "A")) {
		s.Tell(err.Error(), 400)
		return
	}

	user := User{}
	user.FromToken(s.Token)
	user.Name = data.Name
	user.Picture = data.Picture
	err = user.Update()

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(user, 200)

}

// UpdateRoleAPI is the api to update name and/or picture
func UpdateRoleAPI(s auth.Service) {

	var data struct {
		Email    string `json:"email"`
		RoleCode string `json:"role_code"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	user, err := GetUser(data.Email)
	user.RoleCode = data.RoleCode
	err = user.Update()

	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	s.Encode(user, 200)

}
