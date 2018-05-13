package user

var UserTableSQL = `

DROP TABLE IF EXISTS User;
CREATE TABLE User(
	email CHAR(30) PRIMARY KEY NOT NULL,
	name CHAR(20) NOT NULL,
	role_code CHAR(1) DEFAULT 'U'
);

`


