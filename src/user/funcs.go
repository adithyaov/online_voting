package user

import (
	"mysql"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"fmt"
)

func (user *User) Create() (*sql.Result, error) {
	sql, args, err := sq.Insert("User").Columns("name", "email").
						 Values(user.Name, user.Email).ToSql()
	if err != nil {
		return nil, err
	} 
	fmt.Println(sql)
	fmt.Println(args)
	return mysql.Exec(sql, args)
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
	err = rows.Scan(user.Email, user.Name, user.RoleCode)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRoleCode(email string, code string) error {
	return mysql.RunTransaction(mysql.State{ChangeRoleCodeSQL, []interface{}{email, code}})
}



