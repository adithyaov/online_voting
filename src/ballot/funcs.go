package ballot

import (
	//"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	//"github.com/cryptoballot/rsablind"
	"mysql"
	"log"
)

func CreateBallot(id string, name string) Ballot {
	key, err := rsa.GenerateKey(rand.Reader, KeySize)
	if err != nil {
		log.Fatal(err)
	}

	mysql.RunTransaction(mysql.State{
		MakeBallot,
		[]interface{}{id, name, (*(key.N)).String(), (*(key.D)).String(), key.E}})

	return Ballot{id, name, *(key.N), *(key.D), key.E, true}
}

