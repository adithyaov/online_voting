package main

import (
	"net/http"
 	"ballot"
 	"mysql"
)

func main() {
	_ = mysql.RunRawString(ballot.BallotTable)
	b, _ := ballot.CreateBallot("nibbie", "Nib San")
	openBallots := []*ballot.Ballot{b}
	http.HandleFunc("/ballot/create", ballot.CreateAPI)
	http.HandleFunc("/ballot/find", ballot.FindAPI)
	http.HandleFunc("/ballot/delete", ballot.DeleteAPI)
	http.HandleFunc("/ballot/blind", ballot.BlindVoteAPI(openBallots))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}