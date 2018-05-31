package user

// User is type to define a user
type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	RoleCode string `json:"role_code"`
	Picture  string `json:"picture"`
}
