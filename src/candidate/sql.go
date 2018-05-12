package candidate

var CandidateTableSQL = `

DROP TABLE IF EXISTS Candidate;
CREATE TABLE Candidate(
	user_email CHAR(30) NOT NULL,
	ballot_code CHAR(20) NOT NULL,
	details TEXT default "",
	FOREIGN KEY(user_email) REFERENCES User(email) ON DELETE CASCADE,
	FOREIGN KEY(ballot_code) REFERENCES Ballot(code) ON DELETE CASCADE,
	CONSTRAINT unique_foreign UNIQUE (user_email, ballot_code)
);

`


var AddCandidateSQL = `
INSERT INTO Candidate (ballot_code, user_email, details) VALUES (?, ?, "");
`

var UpdateDetailsSQL = `
UPDATE Candidate SET details=? WHERE ballot_code=? AND user_email=?;
`

var RemoveCandidateSQL = `
DELETE FROM Candidate WHERE user_email=?;
`

var GetCandidiateSQL = `
SELECT * FROM Candidate WHERE ballot_code=? AND user_email=?
`


