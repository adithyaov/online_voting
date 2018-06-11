package sql

// Ballots

InsertBallot = `
INSERT INTO Ballot (name, n, d, e, flag) VALUES (?, ?, ?, ?, 1)
`

DeactivateBallot = `
UPDATE Ballot SET flag=0 WHERE id=?
`

BallotName = `
UPDATE Ballot SET name=? WHERE id=?
`

DeleteBallot = `
DELETE Ballot WHERE id=?
`

// Candidates

AddCandidate = `
INSERT INTO Candidates (ballot_id, user_id, details) VALUES (?, ?, "")
`

UpdateDetails = `
UPDATE Candidates SET details=? WHERE ballot_id=? AND user_id=?
`

// Users

RegisterUser = `
INSERT INTO User (name, email) VALUES (?, ?)
`

AddDefRole = `
INSERT INTO UserRoles (role_id, user_id) VALUES (?, ?)
`










