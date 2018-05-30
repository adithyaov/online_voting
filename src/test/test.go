package test

import (
	"auth"
	"ballot"
	"candidate"
	"fmt"
	"mysql"
	"user"
)

func chkErr(id int, err error) {
	if err != nil {
		fmt.Println(id, "----", err.Error())
	}
}

// Auth gives some basic Auth tests
func Auth() {
	err := mysql.RunRawString(user.UserTableSQL)
	chkErr(0, err)

	jwtToken, err := auth.GenerateToken("some google token")
	chkErr(1, err)

	gt, err := auth.ParseToken(jwtToken)
	chkErr(2, err)

	jwtToken2, err := gt.ToToken()
	chkErr(3, err)

	gt2, err := auth.ParseToken(jwtToken2)
	chkErr(4, err)

	fmt.Println(gt == gt2)
	fmt.Println(jwtToken == jwtToken2)

}

// Ballot gives some basic Ballot tests
func Ballot() {
	fmt.Println("------------------BALLOT-------------------")
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

// Candidate gives some basic Candidate tests
func Candidate() {
	fmt.Println("------------------CANDIDATE-------------------")
	err := mysql.RunRawString(candidate.CandidateTableSQL)
	chkErr(199, err)
	b, err := ballot.CreateBallot("a", "A - Ballot")
	chkErr(200, err)
	b.RegexpCandidate = "^(a.*@xx\\.com)$"
	u := user.User{Email: "a@xx.com", Name: "AAADD", RoleCode: "A", Picture: "NOPIC"}
	c := candidate.Candidate{User: &u, Ballot: b}
	fmt.Println(c)
	err = c.Create()
	chkErr(202, err)
	c.Details = "laallalala"
	err = c.UpdateDetails()
	chkErr(203, err)

	c.Nominee1.Valid = true
	c.Nominee1.String = "a@xx.com"
	err = c.UpdateNominees()
	chkErr(204, err)

	c.Nominee1.Valid = true
	c.Nominee1.String = "aa@xx.com"
	err = c.UpdateNominees()
	chkErr(205, err)

	c.Nominee1.Valid = true
	c.Nominee1.String = "aa@xx.com"
	c.Nominee2.Valid = true
	c.Nominee2.String = "aa@xx.com"
	err = c.UpdateNominees()
	chkErr(206, err)

	c2, err := candidate.GetCandidate("a", "a@xx.com")
	chkErr(207, err)
	fmt.Println(c2)
	fmt.Println(c2.User)
	fmt.Println(c2.Ballot)

}
