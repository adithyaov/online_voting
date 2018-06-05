package ballot

import (
	"auth"
	c "common"
	"crypto"
	"crypto/rand"
	rsa "crypto/rsa"
	_ "crypto/sha256" // Required for the init(), registering the hash.
	"encoding/json"
	"errors"
	"fmt"
	"mysql"

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
		Where(sq.Eq{"code": code}).ToSql()

	if err != nil {
		return nil, err
	}

	var ballot Ballot
	var n, d string
	err = mysql.QueryOne(query, args, []interface{}{&ballot.Code, &ballot.Name,
		&n, &d, &ballot.E, &ballot.RegexpVoter,
		&ballot.RegexpCandidate, &ballot.Phase})

	if err != nil {
		return nil, err
	}

	if _, chk := ballot.N.SetString(n, 10); chk != true {
		return nil, errors.New("Could'nt set N as big.Int")
	}
	if _, chk := ballot.D.SetString(d, 10); chk != true {
		return nil, errors.New("Could'nt set D as big.Int")
	}

	return &ballot, nil
}

// GetBallots searches map of openBallots corresponding to the voter Email.
func GetBallots(voterEmail string, openBallots map[string]*Ballot) []*Ballot {
	var ballots []*Ballot

	for _, ballot := range openBallots {
		if c.RegexpStr(ballot.RegexpVoter, voterEmail) == nil {
			ballots = append(ballots, ballot)
		}
	}

	return ballots
}

// DeleteBallot deletes the Ballot from the database
func DeleteBallot(code string) error {

	query, args, err := sq.Delete("Ballot").
		Where(sq.Eq{"code": code}).ToSql()

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
		Where(sq.Eq{"code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// UpdateRegexpCandidate updates candidate regex of corresponding ballot
func (ballot *Ballot) UpdateRegexpCandidate(regexp string) error {
	query, args, err := sq.Update("Ballot").Set("regexp_candidate", regexp).
		Where(sq.Eq{"code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// UpdateName updates ballot name of corresponding ballot
func (ballot *Ballot) UpdateName(name string) error {
	query, args, err := sq.Update("Ballot").Set("name", name).
		Where(sq.Eq{"code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// UpdatePhase updates phase of corresponding ballot
func (ballot *Ballot) UpdatePhase(phase string) error {
	query, args, err := sq.Update("Ballot").Set("phase", phase).
		Where(sq.Eq{"code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// Update updates the prev values to current values in DB
func (ballot *Ballot) Update() error {
	query, args, err := sq.Update("Ballot").Set("phase", ballot.Phase).
		Set("name", ballot.Name).Set("regexp_voter", ballot.RegexpVoter).
		Set("regexp_candidate", ballot.RegexpCandidate).
		Where(sq.Eq{"code": ballot.Code}).ToSql()

	if err != nil {
		return err
	}

	_, err = mysql.Exec(query, args)

	return err
}

// IsCandidate checks if the given candidate is for the specific ballot or not
func (ballot *Ballot) IsCandidate(candidateEmail string) (bool, error) {

	var count int
	query, args, err := sq.Select("COUNT(*)").From("Candidate").
		Where(sq.And{
			sq.Eq{"user_email": candidateEmail},
			sq.Eq{"ballot_code": ballot.Code}}).ToSql()

	if err != nil {
		return false, err
	}

	err = mysql.QueryOne(query, args, []interface{}{&count})
	if err != nil {
		return false, err
	}

	fmt.Println(count)

	if count == 0 {
		return false, nil
	}

	return true, nil
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

// RestartOpenBallotsRT closes all the ballots and reopens them from DB
func RestartOpenBallotsRT(openBallots map[string]*Ballot) error {
	for code := range openBallots {
		delete(openBallots, code)
	}
	// read from db and open all the ballots again
	query, args, err := sq.Select("*").From("Ballot").
		Where(sq.NotEq{"phase": "D"}).ToSql()

	if err != nil {
		return err
	}

	db, err := mysql.OpenDB()

	if err != nil {
		return err
	}

	defer db.Close()

	rows, err := db.Query(query, args...)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var ballot Ballot
		var n, d string
		err := rows.Scan(&ballot.Code, &ballot.Name,
			&n, &d, &ballot.E, &ballot.RegexpVoter,
			&ballot.RegexpCandidate, &ballot.Phase)
		if err != nil {
			return err
		}
		if _, chk := ballot.N.SetString(n, 10); chk != true {
			continue
		}
		if _, chk := ballot.D.SetString(d, 10); chk != true {
			continue
		}
		openBallots[ballot.Code] = &ballot
	}

	return nil
}

// BodyBallotWrapper wraps the functions which require ballot, searches the ballot and
// runs the corresponding function
func BodyBallotWrapper(openBallots map[string]*Ballot, fn func(Service)) func(auth.Service) {
	return func(s auth.Service) {
		var ballotService Service
		ballotService.Service = s
		err := ballotService.FillBallot(openBallots)
		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}

		fn(ballotService)
	}
}

// OpenBallotsWrapper is wrapper over services which need openBallots
func OpenBallotsWrapper(openBallots map[string]*Ballot, fn func(Service)) func(auth.Service) {
	return func(s auth.Service) {
		var ballotService Service
		ballotService.Service = s

		err := ballotService.FillOpenBallots(openBallots)
		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}

		fn(ballotService)
	}
}

// FillBallot fills the Service with the proper ballot
func (s *Service) FillBallot(openBallots map[string]*Ballot) error {
	var data struct {
		BallotCode string `json:"ballot_code"`
	}

	err := json.Unmarshal(s.Body, &data)
	if err != nil {
		return err
	}

	ballot, ok := openBallots[data.BallotCode]
	if ok == false {
		return errors.New("Ballot not found")
	}

	s.Ballot = ballot
	return nil
}

// FillOpenBallots fills the Service with the OpenBallots
func (s *Service) FillOpenBallots(openBallots map[string]*Ballot) error {
	s.OpenBallots = openBallots
	return nil
}
