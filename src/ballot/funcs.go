package ballot

import (
	c "common"
	"crypto"
	"crypto/rand"
	rsa "crypto/rsa"
	_ "crypto/sha256" // Required for the init(), registering the hash.
	"encoding/json"
	"mysql"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/cryptoballot/fdh"
	"github.com/cryptoballot/rsablind"
)

// CreateBallot generates a priv, pub keys and returns
// a *Ballot depending on input params.
func CreateBallot(code string, name string) (*Ballot, error) {
	key, err := rsa.GenerateKey(rand.Reader, KeySize)
	if err != nil {
		return nil, err
	}

	query, args, err := sq.Insert("Ballot").Columns("code", "name", "n", "d", "e").
		Values(code, name, (*(key.PublicKey.N)).String(),
			(*(key.D)).String(), key.PublicKey.E).ToSql()

	if err != nil {
		return nil, err
	}

	_, err = mysql.Exec(query, args)

	if err != nil {
		return nil, err
	}

	ballot := Ballot{code, name, *(key.N), *(key.D), key.E, "^(.*)$", "^(.*)$", "C"}
	return &ballot, nil
}

// OpenBallot searches the database and returns *Ballot
func OpenBallot(code string) (*Ballot, error) {

	query, args, err := sq.Select("*").From("Ballot").
		Where(sq.Eq{"ballot_code": code}).ToSql()

	if err != nil {
		return nil, err
	}

	var ballot Ballot
	var n, d string
	err = mysql.QueryOne(query, args, []interface{}{&ballot.Code, &ballot.Name, &ballot,
		&n, &d, &ballot.E, &ballot.RegexpVoter,
		&ballot.RegexpCandidate, &ballot.Phase})

	if err != nil {
		return nil, err
	}

	if _, chk := ballot.N.SetString(n, 10); chk != true {
		return nil, err
	}
	if _, chk := ballot.D.SetString(d, 10); chk != true {
		return nil, err
	}

	return &ballot, nil
}

// DeleteBallot deletes the Ballot from the database
func DeleteBallot(code string) error {

	query, args, err := sq.Delete("Ballot").
		Where(sq.Eq{"ballot_code": code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err

}

// Hash hashes the Vote type
func (vote *Vote) Hash() ([]byte, error) {
	message, err := json.Marshal(vote)
	if err != nil {
		return nil, err
	}
	return fdh.Sum(crypto.SHA256, HashSize, message), nil
}

// BlindVote take a Vote and returns blinded vote hash and corresponding unblinder
func (ballot *Ballot) BlindVote(vote Vote) ([]byte, []byte, error) {
	publicKey := rsa.PublicKey{N: &(ballot.N), E: ballot.E}
	hashed, err := vote.Hash()
	if err != nil {
		return nil, nil, err
	}

	return rsablind.Blind(&publicKey, hashed)
}

// AddVoter inserts voter into DB corresponding to the given ballot
// to avoid spam.
func (ballot *Ballot) AddVoter(email string) error {

	//Check if valid voter, regexp
	if err := c.RegexpStr(ballot.RegexpVoter, email); err != nil {
		return err
	}

	query, args, err := sq.Insert("BallotUser").Columns("ballot_code", "user_email").
		Values(ballot.Code, email).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// UpdateRegexpVoter updates voter regex of corresponding ballot
func (ballot *Ballot) UpdateRegexpVoter(regexp string) error {
	query, args, err := sq.Update("Ballot").Set("regexp_voter", regexp).
		Where(sq.Eq{"ballot_code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// UpdateRegexpCandidate updates candidate regex of corresponding ballot
func (ballot *Ballot) UpdateRegexpCandidate(regexp string) error {
	query, args, err := sq.Update("Ballot").Set("regexp_candidate", regexp).
		Where(sq.Eq{"ballot_code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// SignBlindHash signs the blinded vote hash and returns the signed hash
func (ballot *Ballot) SignBlindHash(blinded []byte) ([]byte, error) {
	publicKey := rsa.PublicKey{N: &(ballot.N), E: ballot.E}
	privateKey := rsa.PrivateKey{PublicKey: publicKey, D: &(ballot.D)}

	return rsablind.BlindSign(&privateKey, blinded)
}

// UnblindSignedHash takes the signed blind hash and unblinder and returns the
// signed hash.
func (ballot *Ballot) UnblindSignedHash(sign []byte, unblinder []byte) []byte {
	publicKey := rsa.PublicKey{N: &(ballot.N), E: ballot.E}
	return rsablind.Unblind(&publicKey, sign, unblinder)
}

// VerifySign takes the hash and signed hash verifies for the corresponding ballot.
func (ballot *Ballot) VerifySign(hashed []byte, unblindedSign []byte) error {
	publicKey := rsa.PublicKey{N: &(ballot.N), E: ballot.E}
	return rsablind.VerifyBlindSignature(&publicKey, hashed, unblindedSign)
}

// Need to edit this function
func SearchBallotRT(openBallots map[string]*Ballot, ballotCode string) *Ballot {
	if ballot, ok := openBallots[ballotCode]; ok {
		return ballot
	}
	return nil
}

// // Need to edit this function
// func CloseBallotRT(openBallots *([]*Ballot), ballotCode string) {
// 	for i, b := range *openBallots {
// 		if b.Code == ballotCode {
// 			*openBallots = append((*openBallots)[:i], (*openBallots)[i+1:]...)
// 		}
// 	}
// }

// // Need to edit this function
// func OpenBallotRT(openBallots *([]*Ballot), ballotCode string) error {
// 	ballot := SearchBallotRT(openBallots, ballotCode)
// 	if ballot != nil {
// 		return nil
// 	}
// 	ballot, err := OpenBallot(ballotCode)
// 	if err != nil {
// 		return err
// 	}
// 	*openBallots = append(*openBallots, ballot)
// 	return nil
// }

// Wrapper wraps the functions which require ballot, searches the ballot and
// runs the corresponding function
func Wrapper(fn func(http.ResponseWriter, *http.Request, *Ballot, *[]byte), openBallots map[string]*Ballot) c.BodyExtracted {
	return func(w http.ResponseWriter, r *http.Request, body *[]byte) {
		var data struct {
			BallotCode string `json:"ballot_code"`
		}

		err := json.Unmarshal(*body, &data)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		ballot := SearchBallotRT(openBallots, data.BallotCode)
		if ballot == nil {
			http.Error(w, "Ballot not found", 400)
			return
		}
		fn(w, r, ballot, body)
	}
}
