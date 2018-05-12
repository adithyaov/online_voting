package candidate

import (
	"mysql"
	sq "github.com/Masterminds/squirrel"
	"fmt"
	"database/sql"
	"user"
	"ballot"
)

func (candidate *Candidate) Create() (*sql.Result, error) {
	sql, args, err := sq.Insert("Candidate").Columns("user_email", "ballot_code").
					     Values(candidate.User.Email, candidate.Ballot.Code).ToSql()
	if err != nil {
		return nil, err
	} 
	fmt.Println(sql)
	fmt.Println(args)
	return mysql.Exec(sql, args)
}

func (candidate *Candidate) UpdateDetails() (*sql.Result, error) {
	sql, args, err := sq.Update("Candidate").Set("details", candidate.Details).
						 Where(sq.And{
							 sq.Eq{"ballot_code": candidate.Ballot.Code},
							 sq.Eq{"candidate_email": candidate.User.Email}}).ToSql()
	if err != nil {
		return nil, err
	} 
	fmt.Println(sql)
	return mysql.Exec(sql, args)
}

func GetCandidate(code string, email string) (*Candidate, error) {
	sql, args, err := sq.Select("U.name, U.email, B.code, B.name, B.e, B.n, C.details").From("Candidate as C").
						 Join("User as U on U.email = C.user_email").
						 Join("Ballot as B on B.code = C.ballot_code").
						 Where(sq.And{
							 sq.Eq{"C.ballot_code": code},
							 sq.Eq{"C.user_email": email}}).ToSql()
	if err != nil {
		return nil, err
	}
	fmt.Println(sql, args)
	u := user.User{}
	b := ballot.Ballot{}
	c := Candidate{}
	c.User = &u
	c.Ballot = &b
	temp_n := ""
	b.N.SetString(temp_n, 10)
	err = mysql.QueryOne(sql, args, []interface{}{&(u.Name), &(u.Email), &(b.Code), &(b.Name), &(b.E), &(temp_n), &(c.Details)})	
	if err != nil {
		return nil, err
	}
	return &c, err
}





