package ballot

// BallotTable consists of basic SQL statements for init in DB
var BallotTable = `

DROP TABLE IF EXISTS Ballot;
CREATE TABLE Ballot(
	code CHAR(20) PRIMARY KEY,
	name CHAR(40) NOT NULL,
	n TEXT NOT NULL,
	d TEXT NOT NULL,
	e INT NOT NULL,
	regexp_voter TEXT DEFAULT "^(.*)$"
	regexp_candidate TEXT DEFAULT "^(.*)$",
	phase CHAR(1) DEFAULT "C"
);
DROP TABLE IF EXISTS BallotUser;
CREATE TABLE BallotUser(
	user_email CHAR(30) NOT NULL,
	ballot_code CHAR(20) NOT NULL,
	token TEXT DEFAULT "",
	FOREIGN KEY(user_email) REFERENCES User(email) ON DELETE CASCADE,
	FOREIGN KEY(ballot_code) REFERENCES Ballot(code) ON DELETE CASCADE,
	CONSTRAINT unique_foreign UNIQUE (user_email, ballot_code)
);

`
