package candidate

import (
	"ballot"
	c "common"
	"errors"
	"fmt"
	"mysql"
	"user"

	sq "github.com/Masterminds/squirrel"
)

// Create is a function to create candididates
func (candidate *Candidate) Create() error {
	// Check if valid candidate
	if err := c.RegexpStr(candidate.Ballot.RegexpCandidate, candidate.User.Email); err != nil {
		return err
	}

	query, args, err := sq.Insert("Candidate").Columns("user_email", "ballot_code", "details").
		Values(candidate.User.Email, candidate.Ballot.Code, candidate.Details).ToSql()
	if err != nil {
		return err
	}
	fmt.Println(query)
	fmt.Println(args)
	_, err = mysql.Exec(query, args)
	return err
}

// UpdateDetails provides a way to update details
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

// AddNominee adds a nominee to the corresponding candidate
func (candidate *Candidate) AddNominee(nomineeEmail string) error {

	if nomineeEmail == candidate.User.Email {
		return errors.New("Cannot nominate yourself :-|")
	}

	if nomineeEmail == candidate.Nominee1.String ||
		nomineeEmail == candidate.Nominee2.String {
		return errors.New("Already nominated candidate :-|")
	}

	if err := c.RegexpStr(candidate.Ballot.RegexpVoter, nomineeEmail); err != nil {
		return err
	}

	var setField string

	if !candidate.Nominee1.Valid {
		setField = "nominee1_email"
	} else if !candidate.Nominee1.Valid {
		setField = "nominee2_email"
	} else {
		return nil
	}

	query, args, err := sq.Update("Candidate").Set(setField, nomineeEmail).
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

// GetCandidate returns *Candidate after looking up the DB
func GetCandidate(code string, email string) (*Candidate, error) {
	query, args, err := sq.Select("U.name, U.email, B.code, B.name, B.e, B.n, C.details, C.nominee1, C.nominee2").
		From("Candidate as C").
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
	N := ""
	if _, chk := b.N.SetString(N, 10); chk != true {
		return nil, err
	}
	err = mysql.QueryOne(query, args, []interface{}{
		&(u.Name), &(u.Email), &(b.Code),
		&(b.Name), &(b.E), &(N), &(c.Details),
		&(c.Nominee1), &(c.Nominee2)})
	if err != nil {
		return nil, err
	}
	return &c, err
}
