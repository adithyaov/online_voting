package ballot

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"github.com/cryptoballot/rsablind"
	"github.com/cryptoballot/fdh"
	"mysql"
	"encoding/json"
	"net/http"
	c "common"
	sq "github.com/Masterminds/squirrel"
)

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

func DeleteBallot(code string) error {

	query, args, err := sq.Delete("Ballot").
					       Where(sq.Eq{"ballot_code": code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err

}


func (vote *Vote) Hash() ([]byte, error) {
	message, err := json.Marshal(vote)
	if err != nil {
		return nil, err
	}
	return fdh.Sum(crypto.SHA256, HashSize, message), nil
}

func (ballot *Ballot) BlindVote(vote Vote) ([]byte, []byte, error) {
	publicKey := rsa.PublicKey{&(ballot.N), ballot.E}
	hashed, err := vote.Hash()
	if err != nil {
		return nil, nil, err
	}

	return rsablind.Blind(&publicKey, hashed)
}


func (ballot *Ballot) AddVoter(email string) error {

	/*
	Check if valid voter, regexp
	*/
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

func (ballot *Ballot) UpdateRegexpVoter(regexp string) error {
	query, args, err := sq.Update("Ballot").Set("regexp_voter", regexp).
					       Where(sq.Eq{"ballot_code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

func (ballot *Ballot) UpdateRegexpCandidate(regexp string) error {
	query, args, err := sq.Update("Ballot").Set("regexp_candidate", regexp).
					       Where(sq.Eq{"ballot_code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}


func (ballot *Ballot) SignBlindHash(blinded []byte) ([]byte, error) {
	publicKey := rsa.PublicKey{&(ballot.N), ballot.E}
	privateKey := rsa.PrivateKey{PublicKey: publicKey, D: &(ballot.D)}
	
	return rsablind.BlindSign(&privateKey, blinded)
}


func (ballot *Ballot) UnblindSignedHash(sign []byte, unblinder []byte) []byte {
	publicKey := rsa.PublicKey{&(ballot.N), ballot.E}
	return rsablind.Unblind(&publicKey, sign, unblinder)
}

func (ballot *Ballot) VerifySign(hashed []byte, unblindedSign []byte) error {
	publicKey := rsa.PublicKey{&(ballot.N), ballot.E}
	return rsablind.VerifyBlindSignature(&publicKey, hashed, unblindedSign)
}



func SearchBallotRT(openBallots *([]*Ballot), ballotCode string) *Ballot {
	for _, b := range *openBallots {
		if b.Code == ballotCode {
			return b
		}
	}
	return nil
}


func CloseBallotRT(openBallots *([]*Ballot), ballotCode string) {
	for i, b := range *openBallots {
		if b.Code == ballotCode {
			*openBallots = append((*openBallots)[:i], (*openBallots)[i+1:]...)
		}
	}
}

func OpenBallotRT(openBallots *([]*Ballot), ballotCode string) error {
	ballot := SearchBallotRT(openBallots, ballotCode)
	if ballot != nil {
		return nil
	}
	ballot, err := OpenBallot(ballotCode)
	if err != nil {
		return err
	}
	*openBallots = append(*openBallots, ballot)
	return nil
}



func BallotWrapper(fn func(http.ResponseWriter, *http.Request, *Ballot, *[]byte), openBallots *([]*Ballot)) c.BodyExtracted {
	return func (w http.ResponseWriter, r *http.Request, body *[]byte) {
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

