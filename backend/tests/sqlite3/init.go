package sqlite3


import (
    "database/sql"
    "fmt"
    "strconv"
)


func InitDB() {
    fmt.Println(BallotSchema)
    return
}

func main() {
	InitDB()
}

