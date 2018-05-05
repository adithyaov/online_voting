package main

import (
	"fmt"
	"ballot"
	"mysql"
)

func main() {
	fmt.Println("Hi testing! :-)")
	mysql.RunRawString(ballot.BallotTable)
	//mysql.RunTransaction(mysql.State{ballot.MakeBallot, []interface{}{"t1", "test", 5, 20, 22}})
	b := ballot.CreateBallot("elec2019", "S body")
	fmt.Println(b.N.String())
	fmt.Println(b.D.String())
	fmt.Println(b.E)
}