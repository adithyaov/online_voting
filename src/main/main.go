package main

import (
	"ballot"
	c "common"
	"net/http"
)

func main() {
	var openBallots map[string]*ballot.Ballot

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := ballot.Service{Service: c.Service{Writer: w, Request: r},
			Ballot: nil, OpenBallots: nil}
		err := s.FillBallot(openBallots)
		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}
		err = s.FillBody()
		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}
		ballot.CreateAPI(s)
	})
}
