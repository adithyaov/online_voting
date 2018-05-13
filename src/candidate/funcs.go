package candidate

import (
	"mysql"
	sq "github.com/Masterminds/squirrel"
	"fmt"
	"user"
	"ballot"
)

func (candidate *Candidate) Create() error {
	/*
	Check if valid candidate
	*/
	query, args, err := sq.Insert("Candidate").Columns("user_email", "ballot_code").
					       Values(candidate.User.Email, candidate.Ballot.Code).ToSql()
	if err != nil {
		return err
	} 
	fmt.Println(query)
	fmt.Println(args)
	_, err = mysql.Exec(query, args)
	return err
}

func (candidate *Candidate) UpdateDetails() error {
	query, args, err := sq.Update("Candidate").Set("details", candidate.Details).
						   Where(sq.And{
							   sq.Eq{"ballot_code": candidate.Ballot.Code},
							   sq.Eq{"candidate_email": candidate.User.Email}}).ToSql()
	if err != nil {
		return err
	} 
	fmt.Println(query)
	_, err = mysql.Exec(query, args)
	return err
}

func GetCandidate(code string, email string) (*Candidate, error) {
	query, args, err := sq.Select("U.name, U.email, B.code, B.name, B.e, B.n, C.details, C.nominee1, C.nominee2").From("Candidate as C").
						   Join("User as U on U.email = C.user_email").
						   Join("Ballot as B on B.code = C.ballot_code").
						   Where(sq.And{
							   sq.Eq{"C.ballot_code": code},
							   sq.Eq{"C.user_email": email}}).ToSql()
	if err != nil {
		return nil, err
	}
	fmt.Println(query, args)
	u := user.User{}
	b := ballot.Ballot{}
	c := Candidate{}
	c.User = &u
	c.Ballot = &b
	temp_n := ""
	if _, chk := b.N.SetString(temp_n, 10); chk != true {
		return nil, err
	}
	err = mysql.QueryOne(query, args, []interface{}{&(u.Name), &(u.Email), &(b.Code), 
													&(b.Name), &(b.E), &(temp_n), &(c.Details),
													&(c.Nominee1), &(c.Nominee2)})	
	if err != nil {
		return nil, err
	}
	return &c, err
}





