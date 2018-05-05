package main

import (
	"fmt"
	"mysql"
)

func main() {
	r, _ := mysql.RunQuery(mysql.State{"SELECT * FROM Ballot WHERE code=?", []interface{}{"elec2019"}})
	var code, name, n, d string
	var e int
	var flag bool
	for r.Next() {
		_ = r.Scan(&(code), &(name), &(n), &(d), &(e), &(flag))
		fmt.Println(code, name, n, d, e, flag)
	}
}