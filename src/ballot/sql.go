package ballot

var MakeBallot = `
INSERT INTO Ballot (code, name, n, d, e, flag) VALUES (?, ?, ?, ?, ?, 1)
`

var DeactivateBallot = `
UPDATE Ballot SET flag=0 WHERE code=?
`

var BallotName = `
UPDATE Ballot SET name=? WHERE code=?
`

var DeleteBallot = `
DELETE Ballot WHERE code=?
`

var GetBallot = `
SELECT * FROM Ballot WHERE code=?
`

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
