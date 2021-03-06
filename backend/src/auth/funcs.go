package auth

import (
	c "common"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ToClaims converts a GoogleToken into jwt.MapClaims.
func (gt GoogleToken) ToClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"email":          gt.Email,
		"email_verified": strconv.FormatBool(gt.EmailVerified),
		"iat":            strconv.FormatInt(gt.Iat.Unix(), 10),
		"exp":            strconv.FormatInt(gt.Exp.Unix(), 10),
		"name":           gt.Name,
		"aud":            gt.Aud,
		"role_code":      gt.RoleCode,
		"picture":        gt.Picture,
	}
}

// ToToken converts GoogleToken to a jwt token string.
func (gt GoogleToken) ToToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, gt.ToClaims())
	tokenString, err := token.SignedString(SessionSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateToken generates a jwt token from the token given
// by google plus login button.
// https://developers.google.com/identity/sign-in/web/backend-auth
func GenerateToken(googleToken string) (string, error) {
	// Verify google token

	// Dummy Vals
	email := "a@gmail.com"
	name := "A"
	picture := "pic1"
	rc := "AMUX"

	// moved user addition logic to user
	/*
		1. First Populate user with token
		2. check if user exists
		3. Create User if does not exist
	*/

	// Also put nbf in the following claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":          email,
		"email_verified": "true",
		"iat":            strconv.FormatInt(time.Now().Unix(), 10),
		"exp":            strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10),
		"name":           name,
		"aud":            "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com",
		"role_code":      rc,
		"picture":        picture,
	})
	tokenString, err := token.SignedString(SessionSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses the jwt token and returns a GoogleToken.
func ParseToken(tokenString string) (GoogleToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SessionSecret, nil
	})
	if err != nil {
		return GoogleToken{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Create a proper google token
		iat, err := strconv.ParseInt(claims["iat"].(string), 10, 64)
		if err != nil {
			return GoogleToken{}, err
		}
		exp, err := strconv.ParseInt(claims["exp"].(string), 10, 64)
		if err != nil {
			return GoogleToken{}, err
		}
		emailVer, err := strconv.ParseBool(claims["email_verified"].(string))
		if err != nil {
			return GoogleToken{}, err
		}
		return GoogleToken{claims["aud"].(string), time.Unix(iat, 0), time.Unix(exp, 0),
			claims["email"].(string), claims["role_code"].(string), emailVer,
			claims["name"].(string), claims["picture"].(string)}, nil
	}
	return GoogleToken{}, fmt.Errorf("Invalid Token")
}

// Wrapper wraps the corresponding function after checking
// the jwt token in the header
func Wrapper(validCodeRune string, fn func(Service)) func(c.Service) {
	return func(sBasic c.Service) {

		s := Service{}
		s.Service = sBasic

		err := s.ExtractToken()

		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}

		err = errors.New("Permission denied")
		if c.IsIn(s.Token.RoleCode, validCodeRune) {
			err = nil
		}

		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}

		fn(s)
	}
}

// ExtractToken extracts the token from the header and sets it in the service
func (s *Service) ExtractToken() error {
	token := s.Request.Header["Token"]
	var gt GoogleToken
	var err error
	err = nil

	if len(token) == 0 {
		gt = GoogleToken{RoleCode: "X"}
	} else {
		gt, err = ParseToken(token[0])
	}

	if err != nil {
		return err
	}

	s.Token = gt
	return nil

}

// IsOwnerOr gived bool for the fact that he is an owner or a part of code rune
func IsOwnerOr(emailRecieved string, gt GoogleToken, validCodeRune string) bool {

	if emailRecieved == gt.Email || c.IsIn(gt.RoleCode, validCodeRune) {
		return true
	}

	return false
}
