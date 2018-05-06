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
)

func CreateBallot(code string, name string) (*Ballot, error) {
	key, err := rsa.GenerateKey(rand.Reader, KeySize)
	if err != nil {
		return nil, err
	}

	err = mysql.RunTransaction(mysql.State{
			MakeBallot,
			[]interface{}{code, name, (*(key.PublicKey.N)).String(), (*(key.D)).String(), key.PublicKey.E}})
	if err != nil {
		return nil, err
	}
	
	ballot := Ballot{code, name, *(key.N), *(key.D), key.E, true}
	return &ballot, nil
}

func OpenBallot(code string) (*Ballot, error) {
	rows, err := mysql.RunQuery(mysql.State{GetBallot, []interface{}{code}})
	if err != nil {
		return nil, err
	}
	
	rows.Next()

	var ballot Ballot
	var n, d string
	rows.Scan(&ballot.Code, &ballot.Name, &ballot, &n, &d, &ballot.E, &ballot.Flag)
	ballot.N.SetString(n, 10)
	ballot.D.SetString(d, 10)
	return &ballot, nil
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










