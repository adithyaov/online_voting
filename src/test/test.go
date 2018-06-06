package test

import (
	"auth"
	"ballot"
	"candidate"
	"database/sql"
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
	err := mysql.RunRawString(ballot.BallotTableSQL)
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

// User gives some basic User tests
func User() {
	err := mysql.RunRawString(user.UserTableSQL)
	chkErr(500, err)

	jwtToken, err := auth.GenerateToken("some google token")
	chkErr(501, err)

	gt, err := auth.ParseToken(jwtToken)
	chkErr(502, err)

	u := user.User{}
	u.FromToken(gt)
	exists, err := u.CheckIfExists()
	chkErr(503, err)
	fmt.Println(exists)

	err = u.Create()
	chkErr(504, err)
	fmt.Println(u)

	err = u.Create()
	chkErr(505, err)

	exists, err = u.CheckIfExists()
	chkErr(506, err)
	fmt.Println(exists)

	u2 := user.User{}
	err = u2.SetWith("a@xx.com")
	chkErr(509, err)
	fmt.Println(u)
	fmt.Println(u2)
	fmt.Println(u == u2)

	u.Name = "LALALA"
	err = u.Update()
	chkErr(510, err)

	err = user.DeleteUser(u.Email)
	chkErr(507, err)

	exists, err = u.CheckIfExists()
	chkErr(508, err)
	fmt.Println(exists)
}

// Init inits all the DB related and sets up dummy info.
func Init() {
	err := mysql.RunRawString(user.UserTableSQL)
	chkErr(0, err)
	err = mysql.RunRawString(ballot.BallotTableSQL)
	chkErr(1, err)
	err = mysql.RunRawString(candidate.CandidateTableSQL)
	chkErr(2, err)
	b, err := ballot.CreateBallot("a", "A")
	chkErr(3, err)
	u1 := user.User{Email: "a@gmail.com", Name: "A", RoleCode: "AMUX", Picture: "pic1"}
	u2 := user.User{Email: "b@gmail.com", Name: "B", RoleCode: "AMUX", Picture: "pic2"}
	u3 := user.User{Email: "c@gmail.com", Name: "C", RoleCode: "AMUX", Picture: "pic3"}
	err = u1.Create()
	chkErr(4, err)
	err = u2.Create()
	chkErr(5, err)
	err = u3.Create()
	chkErr(6, err)
	can := candidate.Candidate{User: &u1, Ballot: b,
		Nominee1: sql.NullString{Valid: true, String: u2.Email},
		Nominee2: sql.NullString{Valid: true, String: u3.Email},
		Details:  "none://doc.dtx"}
	err = can.Create()
	chkErr(7, err)
	err = can.UpdateDetails()
	chkErr(8, err)
	err = can.UpdateNominees()
	chkErr(9, err)

}
