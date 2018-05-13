package candidate

var CandidateTableSQL = `

DROP TABLE IF EXISTS Candidate;
CREATE TABLE Candidate(
	user_email CHAR(30) NOT NULL,
	ballot_code CHAR(20) NOT NULL,
	details TEXT DEFAULT,
	nominee1_email CHAR(30),
	nominee2_email CHAR(30),
	FOREIGN KEY(user_email) REFERENCES User(email) ON DELETE CASCADE,
	FOREIGN KEY(ballot_code) REFERENCES Ballot(code) ON DELETE CASCADE,
	CONSTRAINT unique_foreign UNIQUE (user_email, ballot_code)
);

`



