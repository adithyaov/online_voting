package main

import (
	"fmt"
	"ballot"
	"mysql"
)

func main() {
	fmt.Println("Hi testing! :-)")
	_ = mysql.RunRawString(ballot.BallotTable)
	b, _ := ballot.CreateBallot("elec2019", "S body")
	v := ballot.Vote{"a", "b", "c"}
	blinded, unblinder, _ := b.BlindVote(v)
	sig, _ := b.SignBlindHash(blinded)
	us := b.UnblindSignedHash(sig, unblinder)
	hashed, _ := v.Hash()
	fmt.Println(b.VerifySign(hashed, us))
}