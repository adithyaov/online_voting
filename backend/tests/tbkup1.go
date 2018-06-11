package types

import (
	"crypto/rsa"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID uint `gorm:"primary_key"`
	Name string
	Email string
	Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID uint `gorm:"primary_key"`
	RoleName string
}

type RoleMap struct {
	UserID uint
	User User `gorm:"foreignkey:UserID"`
	RoleID uint
	Role Role 
}

type Candidate struct {
	UserID uint
	User User `gorm:"foreignkey:UserID"`
	Poster string
	BallotID uint
	Ballot Ballot `gorm:"foreignkey:BallotID"`
}

type Ballot struct {
	ID uint `gorm:"primary_key"`
	Name string
	N big.Int
	D big.Int
	E big.Int
	flag bool
}


