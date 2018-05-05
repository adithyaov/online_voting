package ballot

var MakeBallot = `
INSERT INTO Ballot (id, name, n, d, e, flag) VALUES (?, ?, ?, ?, ?, 1)
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

var GetBallot = `
SELECT * FROM Ballot WHERE id=?
`

var BallotTable = `

DROP TABLE IF EXISTS Ballot;
CREATE TABLE Ballot(
	id CHAR(20) PRIMARY KEY,
	name CHAR(40) NOT NULL,
	n TEXT NOT NULL,
	d TEXT NOT NULL,
	e INT NOT NULL,
	flag BOOL DEFAULT 1
);

`
