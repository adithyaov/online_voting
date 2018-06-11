package main

import (
    "fmt"
    "encoding/json"
)

type Another struct {
	A *string
	B string
}

type testS struct {
	I int
	S string
	AA Another
}

// type Employee struct {
//     Name string `json:"empname"`
//     Number int  `json:"empid"`
// }

func main() {
	x := "hi"
	emp := testS{S: "Rocky",I: 5454, AA: Another{A: &x, B: "lalala"}}
    e, err := json.Marshal(emp)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(e))
}

