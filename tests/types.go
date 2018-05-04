package types

import "crypto/rsa"

type User struct {
	ID uint
	Name string
	Email string
	Roles []*Role
}

type Role struct {
	ID uint
	RoleName string
}

type Candidate struct {
	User *User
	PosterUrl string
	Ballot *Ballot
}

type Ballot struct {
	ID uint
	Name string
	N big.Int
	D big.Int
	E int
}


