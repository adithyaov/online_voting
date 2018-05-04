package ballot

var MakeBallot = `
INSERT INTO Ballot (name, n, d, e, flag) VALUES (?, ?, ?, ?, 1)
`

var DeactivateBallot = `
UPDATE Ballot SET flag=0 WHERE id=?
`

var BallotName = `
UPDATE Ballot SET name=? WHERE id=?
`

var DeleteBallot = `
DELETE Ballot WHERE id=?
`

var BallotTable = `

DROP TABLE IF EXISTS Ballot;
CREATE TABLE Ballot(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name CHAR(40) NOT NULL,
	n BIGINT NOT NULL,
	d BIGINT NOT NULL,
	e INT NOT NULL,
	flag BOOL DEFAULT 1
);

`
