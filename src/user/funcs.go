package user

import (
	"mysql"
	sq "github.com/Masterminds/squirrel"
	"fmt"
)

func (user *User) Create() error {
	query, args, err := sq.Insert("User").Columns("name", "email").
						 Values(user.Name, user.Email).ToSql()
	if err != nil {
		return err
	} 
	fmt.Println(query)
	fmt.Println(args)
	_, err = mysql.Exec(query, args)
	return err
}

func DeleteUser(email string) error {
	query, args, err := sq.Delete("User").
						 Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return err
	} 
	_, err = mysql.Exec(query, args)
	return err
}

func (user *User) SetWith(email string) error {
	query, args, err := sq.Select("email, name, role_code").From("User").
						 Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return err
	} 

	err = mysql.QueryOne(query, args, []interface{}{user.Email, user.Name, user.RoleCode})
	if err != nil {
		return err
	}
	return nil
}

func UpdateRoleCode(email string, code string) error {
	query, args, err := sq.Update("User").Set("role_code", code).
						   Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return err
	} 

	_, err = mysql.Exec(query, args)
	if err != nil {
		return err
	}
	return nil
}



