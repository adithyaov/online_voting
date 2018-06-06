package main

import (
	"auth"
	"ballot"
	c "common"
	"net/http"
	"os"
)

func main() {
	openBallots := make(map[string]*ballot.Ballot)
	err := ballot.RestartOpenBallotsRT(openBallots)

	if err != nil {
		os.Exit(1)
	}

	http.HandleFunc("/", c.CreateService(c.BodyCheckWrapper(auth.Wrapper("A", ballot.CreateAPI))))

	http.ListenAndServe(":8080", nil)

}
