package common

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)


func BodyCheckWrapper(fn BodyExtracted) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if len(body) > MaxReqBody {
			http.Error(w, "Body too long :-(", 400)
			return
		}
		fn(w, r, &body)
	}
}


func ConvertBSToIS(bSlice []byte) []int {
	var iSlice []int
	for _, b := range bSlice {
		iSlice = append(iSlice, int(b))
	}
	return iSlice
}

func ConvertISToBS(iSlice []int) []byte {
	var bSlice []byte
	for _, b := range iSlice {
		bSlice = append(bSlice, byte(b))
	}
	return bSlice
}

func RegexpStr(expr string, str string) error {
	matched, err := regexp.MatchString(expr, str)

	if err != nil {
		return err
	}

	if matched != true {
		return fmt.Errorf("Invalid Voter.")
	}

	return nil

}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	    }
	    return SessionSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	    return claims, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}


