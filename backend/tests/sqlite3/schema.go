package sqlite3


// Create tables

BallotSchema = `

CREATE TABLE Ballot(
	id INT PRIMARY KEY NOT NULL,
	name CHAR(40) NOT NULL,
	n BIGINT NOT NULL,
	d BIGINT NOT NULL,
	e INT NOT NULL,
	flag BOOL DEFAULT 1
)

`

UserSchema = `

CREATE TABLE User(
	id INT PRIMARY KEY NOT NULL,
	name CHAR(20) NOT NULL,
	email CHAR(30) NOT NULL
)

`

RoleSchema  = `

CREATE TABLE Role(
	id INT PRIMARY KEY NOT NULL,
	role_name CHAR(20) NOT NULL
)

`

UserRolesSchema = `

CREATE TABLE UserRoles(
	user_id INT NOT NULL,
	role_id INT NOT NULL,
	FOREIGN KEY(user_id) REFERENCES User(id) ON DELETE CASCADE,
	FOREIGN KEY(role_id) REFERENCES Role(id) ON DELETE CASCADE,
	CONSTRAINT unique_foreign UNIQUE (user_id, role_id)
)

`

CandidateSchema = `

CREATE TABLE Candidate(
	user_id INT NOT NULL,
	ballot_id INT NOT NULL,
	details TEXT,
	FOREIGN KEY(user_id) REFERENCES User(id) ON DELETE CASCADE,
	FOREIGN KEY(ballot_id) REFERENCES Ballot(id) ON DELETE CASCADE,
	CONSTRAINT unique_foreign UNIQUE (user_id, ballot_id)
)

`















