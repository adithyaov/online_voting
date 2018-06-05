package main

import (
	"ballot"
	c "common"
	"net/http"
)

func main() {
	openBallots := make(map[string]*ballot.Ballot)

	http.HandleFunc("/", c.CreateService(c.BodyCheckWrapper(
		ballot.BodyBallotWrapper(openBallots, ballot.CreateAPI))))
}
