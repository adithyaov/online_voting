package ballot

var BallotTable = `

DROP TABLE IF EXISTS Ballot;
CREATE TABLE Ballot(
	code CHAR(20) PRIMARY KEY,
	name CHAR(40) NOT NULL,
	n TEXT NOT NULL,
	d TEXT NOT NULL,
	e INT NOT NULL,
	flag BOOL DEFAULT 1
);

`
