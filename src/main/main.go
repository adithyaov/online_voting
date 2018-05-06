package main

import (
	"net/http"
 	"ballot"
 	"mysql"
 	c "common"
)

func main() {
	_ = mysql.RunRawString(ballot.BallotTable)
	b, _ := ballot.CreateBallot("nibbie", "Nib San")
	openBallots := []*ballot.Ballot{b}
	http.HandleFunc("/ballot/create", c.BodyCheckWrapper(ballot.CreateAPI))
	http.HandleFunc("/ballot/find", c.BodyCheckWrapper(ballot.FindAPI))
	http.HandleFunc("/ballot/delete", c.BodyCheckWrapper(ballot.DeleteAPI))
	http.HandleFunc("/ballot/blind", c.BodyCheckWrapper(ballot.BlindVoteAPI(openBallots)))
	http.HandleFunc("/ballot/sign", c.BodyCheckWrapper(ballot.SignBytesAPI(openBallots)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}