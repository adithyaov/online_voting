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
			sq.Eq{"user_email": candidate.User.Email}}).ToSql()
	if err != nil {
		return err
	}
	fmt.Println(query)
	_, err = mysql.Exec(query, args)
	return err
}

// UpdateNominees updates the nominees after basic checks
func (candidate *Candidate) UpdateNominees() error {

	if candidate.Nominee1.String == candidate.User.Email ||
		candidate.Nominee2.String == candidate.User.Email {
		return errors.New("Cannot nominate yourself :-|")
	}

	if candidate.Nominee1.String == candidate.Nominee2.String &&
		candidate.Nominee1.Valid && candidate.Nominee2.Valid {
		return errors.New("Double nomination not permitted :-|")
	}

	if err := c.RegexpStr(candidate.Ballot.RegexpVoter, candidate.Nominee1.String); err != nil && candidate.Nominee1.Valid {
		return err
	}

	if err := c.RegexpStr(candidate.Ballot.RegexpVoter, candidate.Nominee2.String); err != nil && candidate.Nominee2.Valid {
		return err
	}

	builder := sq.Update("Candidate")

	if candidate.Nominee1.Valid {
		builder = builder.Set("nominee1_email", candidate.Nominee1.String)
	}
	if candidate.Nominee2.Valid {
		builder = builder.Set("nominee2_email", candidate.Nominee2.String)
	}

	query, args, err := builder.Where(sq.And{
		sq.Eq{"ballot_code": candidate.Ballot.Code},
		sq.Eq{"user_email": candidate.User.Email}}).ToSql()
	if err != nil {
		return err
	}
	fmt.Println(query)
	_, err = mysql.Exec(query, args)
	return err
}

// GetCandidate returns *Candidate after looking up the DB
func GetCandidate(code string, email string) (*Candidate, error) {
	query, args, err := sq.Select("U.name, U.email, U.picture, B.code, B.name, B.e, B.n, C.details, C.nominee1_email, C.nominee2_email").
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
	n := ""
	err = mysql.QueryOne(query, args, []interface{}{
		&(u.Name), &(u.Email), &(u.Picture), &(b.Code),
		&(b.Name), &(b.E), &(n), &(c.Details),
		&(c.Nominee1), &(c.Nominee2)})
	if err != nil {
		return nil, err
	}

	if _, chk := b.N.SetString(n, 10); chk != true {
		return nil, errors.New("Could'nt set N as big.Int")
	}

	return &c, nil
}
