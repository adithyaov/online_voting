package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ToClaims converts a GoogleToken into jwt.MapClaims.
func (gt GoogleToken) ToClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"email":          gt.Email,
		"email_verified": strconv.FormatBool(gt.EmailVerified),
		"iat":            gt.Iat.Unix(),
		"exp":            gt.Exp.Unix(),
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

	// Also put nbf in the following claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":          "111501017@smail.iitpkd.ac.in",
		"role_code":      "U",
		"email_verified": "true",
		"iat":            "1433978353",
		"exp":            "1433981953",
		"name":           "Adithya O V",
		"picture":        "ssdsds",
		"aud":            "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com",
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
			claims["email"].(string), claims["email"].(string), emailVer,
			claims["name"].(string), claims["picture"].(string)}, nil
	}
	return GoogleToken{}, fmt.Errorf("Invalid Token")
}

// Wrapper wraps the corresponding function after checking
// the jwt token in the header
func Wrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header["token"][0]
		_, err := ParseToken(token)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fn(w, r)
	}
}
