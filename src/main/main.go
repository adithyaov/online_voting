package main

import (
	"auth"
	"ballot"
	"candidate"
	c "common"
	"net/http"
	"os"
	"user"
)

func main() {
	openBallots := make(map[string]*ballot.Ballot)
	err := ballot.RestartOpenBallotsRT(openBallots)

	if err != nil {
		os.Exit(1)
	}

	// Ballot EP's

	http.HandleFunc("/ballot/create", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("A", ballot.CreateAPI)))))

	http.HandleFunc("/ballot/find", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUMX", ballot.FindAPI)))))

	http.HandleFunc("/ballot/delete", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("A", ballot.DeleteAPI)))))

	http.HandleFunc("/ballot/blind-vote", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUMX",
			ballot.BodyBallotWrapper(openBallots, ballot.BlindVoteAPI))))))

	http.HandleFunc("/ballot/sign-bytes", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM",
			ballot.BodyBallotWrapper(openBallots, ballot.SignBytesAPI))))))

	http.HandleFunc("/ballot/unblind-sign", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUMX",
			ballot.BodyBallotWrapper(openBallots, ballot.UnblindSignAPI))))))

	http.HandleFunc("/ballot/verify-sign", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUMX",
			ballot.BodyBallotWrapper(openBallots, ballot.VerifySignAPI))))))

	http.HandleFunc("/ballot/find-ballots", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM",
			ballot.BodyBallotWrapper(openBallots, ballot.FindBallotsAPI))))))

	// Candidate EP's

	http.HandleFunc("/candidate/create", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM", candidate.CreateAPI)))))

	http.HandleFunc("/candidate/add-nominee", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM", candidate.AddNomineeAPI)))))

	http.HandleFunc("/candidate/update-details", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM", candidate.UpdateDetailsAPI)))))

	http.HandleFunc("/candidate/delete", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM", candidate.DeleteAPI)))))

	// User EP's

	http.HandleFunc("/user/token-service", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUMX", user.AuthUserAPI)))))

	http.HandleFunc("/user/delete", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM", user.DeleteAPI)))))

	http.HandleFunc("/user/delete", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("AUM", user.UpdatePersonalAPI)))))

	http.HandleFunc("/user/delete", c.CreateService(c.MethodWrapper("POST",
		c.BodyCheckWrapper(auth.Wrapper("A", user.UpdateRoleAPI)))))

	http.ListenAndServe(":8080", nil)

}
