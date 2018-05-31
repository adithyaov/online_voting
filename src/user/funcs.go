package user

import (
	"auth"
	"fmt"
	"mysql"

	sq "github.com/Masterminds/squirrel"
)

// FromToken creates a user from token recieved
func (user *User) FromToken(gt auth.GoogleToken) {
	user.Email = gt.Email
	user.Name = gt.Name
	user.RoleCode = gt.RoleCode
	user.Picture = gt.Picture
}

func (user *User) CheckIfExists() (bool, error) {

	var count int
	query, args, err := sq.Select("COUNT(*)").From("User").
		Where(sq.Eq{"email": user.Email}).ToSql()
	err = mysql.QueryOne(query, args, []interface{}{&count})
	if err != nil {
		return false, err
	}

	fmt.Println(count)

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// Create is a basic function to create a user
func (user *User) Create() error {
	query, args, err := sq.Insert("User").Columns("name", "email", "role_code", "picture").
		Values(user.Name, user.Email, user.RoleCode, user.Picture).ToSql()
	if err != nil {
		return err
	}
	fmt.Println(query)
	fmt.Println(args)
	_, err = mysql.Exec(query, args)
	return err
}

// DeleteUser is a basic function to delete the user
func DeleteUser(email string) error {
	query, args, err := sq.Delete("User").
		Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return err
	}
	_, err = mysql.Exec(query, args)
	return err
}

// SetWith sets user with the data of specific email
func (user *User) SetWith(email string) error {
	query, args, err := sq.Select("email, name, role_code, picture").From("User").
		Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return err
	}

	err = mysql.QueryOne(query, args, []interface{}{user.Email, user.Name, user.RoleCode, user.Picture})
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoleCode update the role_code of the user
func UpdateRoleCode(email string, roleCode string) error {
	query, args, err := sq.Update("User").Set("role_code", roleCode).
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

// Update updates the user with the current values
func (user *User) Update() error {
	query, args, err := sq.Update("User").Set("role_code", user.RoleCode).
		Set("name", user.Name).
		Set("picture", user.Picture).
		Where(sq.Eq{"email": user.Email}).ToSql()
	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)
	if err != nil {
		return err
	}
	return nil
}
