package user

var UserTableSQL = `

CREATE TABLE User(
	email CHAR(30) PRIMARY KEY NOT NULL,
	name CHAR(20) NOT NULL,
	role_code CHAR(1) DEFAULT 'U'
);

`

var RegisterUserSQL = `
INSERT INTO User (name, email) VALUES (?, ?);
`

var ChangeRoleCodeSQL = `
UPDATE User SET role_code=? WHERE email=?;
`

var SelectUserSQL = `
SELECT * FROM User WHERE email=?
`

var DeleteUserSQL = `
DELETE FROM User WHERE email=?
`



