package test

import (
	"auth"
	"fmt"
)

func chkErr(id int, err error) {
	if err != nil {
		fmt.Println(id, "----", err.Error())
	}
}

// Auth gives some basic Auth tests
func Auth() {
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
