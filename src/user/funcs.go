package user

import (
	"mysql"
)

func (user *User) Create() error {
	return mysql.RunTransaction(mysql.State{RegisterUserSQL, []interface{}{user.Name, user.Email}})
}

func DeleteUser(email string) error {
	return mysql.RunTransaction(mysql.State{DeleteUserSQL, []interface{}{email}})
}

func (user *User) SetWith(email string) error {
	rows, err := mysql.RunQuery(mysql.State{SelectUserSQL, []interface{}{user.Email}})
	if err != nil {
		return err
	}
	rows.Next()
	rows.Scan(user.Email, user.Name, user.RoleCode)
	return nil
}

func UpdateRoleCode(email string, code string) error {
	return mysql.RunTransaction(mysql.State{ChangeRoleCodeSQL, []interface{}{email, code}})
}



