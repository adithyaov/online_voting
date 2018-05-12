package main

import (
	"net/http"
 	"ballot"
 	"mysql"
 	com "common"
 	"candidate"
 	"user"
 	"fmt"
)

func main() {
	err := mysql.RunRawString(ballot.BallotTable)
	fmt.Println(err)
	err = mysql.RunRawString(candidate.CandidateTableSQL)
	fmt.Println(err)
	err = mysql.RunRawString(user.UserTableSQL)
	fmt.Println(err)

	d := candidate.Candidate{}
	u := user.User{Email:"aa@aa.com", Name: "nanana"}
	_, err = u.Create()
	fmt.Println("usr crt", err)
	b, err := ballot.CreateBallot("nibbie", "chubbie")
	fmt.Println("blt crt", err)
	d.User = &u
	d.Ballot = b
	_, err = d.Create()
	fmt.Println("can crt", err)
	c, err := candidate.GetCandidate("nibbie", "aa@aa.com")
	fmt.Println(c, err)
	

	openBallots := []*ballot.Ballot{b}
	http.HandleFunc("/ballot/create", com.BodyCheckWrapper(ballot.CreateAPI))
	http.HandleFunc("/ballot/find", com.BodyCheckWrapper(ballot.FindAPI))
	http.HandleFunc("/ballot/delete", com.BodyCheckWrapper(ballot.DeleteAPI))
	http.HandleFunc("/ballot/blind", com.BodyCheckWrapper(ballot.BallotWrapper(ballot.BlindVoteAPI, &openBallots)))
	http.HandleFunc("/ballot/sign", com.BodyCheckWrapper(ballot.BallotWrapper(ballot.SignBytesAPI, &openBallots)))
	http.HandleFunc("/ballot/unblind", com.BodyCheckWrapper(ballot.BallotWrapper(ballot.UnblindSignAPI, &openBallots)))
	http.HandleFunc("/ballot/verify", com.BodyCheckWrapper(ballot.BallotWrapper(ballot.VerifySignAPI, &openBallots)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}