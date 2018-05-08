package candidate

var CandidateTableSQL = `

DROP TABLE IF EXISTS Candidate;
CREATE TABLE Candidate(
	user_email CHAR(30) NOT NULL,
	ballot_code CHAR(20) NOT NULL,
	details TEXT,
	FOREIGN KEY(user_email) REFERENCES User(email) ON DELETE CASCADE,
	FOREIGN KEY(ballot_code) REFERENCES Ballot(code) ON DELETE CASCADE,
	CONSTRAINT unique_foreign UNIQUE (user_id, ballot_id)
);

`


var AddCandidate = `
INSERT INTO Candidate (ballot_code, user_email, details) VALUES (?, ?, "");
`

var UpdateDetails = `
UPDATE Candidate SET details=? WHERE ballot_code=? AND user_email=?;
`

var RemoveCandidate = `
DELETE FROM Candidate WHERE user_email=?;
`


