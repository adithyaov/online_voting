package test

import (
	"ballot"
	"fmt"
	"mysql"
)

// Ballot gives some basic Ballot tests
func Ballot() {
	err := mysql.RunRawString(ballot.BallotTable)
	chkErr(0, err)
	b, err := ballot.CreateBallot("a", "A - Ballot")
	chkErr(1, err)
	b2, err := ballot.OpenBallot("a")
	chkErr(15, err)
	fmt.Println(b.N)
	fmt.Println(b2.N)
	fmt.Println(b2.E == b.E)
	err = ballot.DeleteBallot("a")
	chkErr(2, err)

}

// Vote gives some basic Vote tests
func Vote() {
	fmt.Println("------------------VOTE-------------------")
	b, err := ballot.CreateBallot("a", "A - Ballot")
	chkErr(100, err)
	v := ballot.Vote{BallotCode: "a",
		CandidateEmail: "adi@gmail.com", Bias: "78egw287e"}
	vHash, err := v.Hash()
	chkErr(101, err)
	fmt.Println(vHash)

	bv, ub, err := b.BlindVote(v)
	chkErr(103, err)
	fmt.Println(bv)
	fmt.Println(ub)

	err = b.AddVoter(v.CandidateEmail)
	chkErr(105, err)
	err = b.AddVoter(v.CandidateEmail)
	chkErr(105, err)

	err = b.UpdateRegexpVoter("l.*@ll\\.com")
	chkErr(111, err)
	err = b.UpdateRegexpCandidate("ll.*@ll\\.com")
	chkErr(112, err)
	err = b.UpdateName("Lala land")
	chkErr(113, err)
	err = b.UpdatePhase("N")
	chkErr(114, err)

	b2, err := ballot.OpenBallot("a")
	chkErr(115, err)
	fmt.Println(b2.Name, b2.RegexpVoter, b2.Phase, b2.RegexpCandidate)

	s, err := b.SignBlindHash(bv)
	chkErr(116, err)

	sh := b.UnblindSignedHash(s, ub)

	err = b.VerifySign(vHash, sh)
	chkErr(117, err)

	var openBallots = make(map[string]*ballot.Ballot)
	err = ballot.RestartOpenBallotsRT(openBallots)
	chkErr(118, err)
	fmt.Println(openBallots)
}
