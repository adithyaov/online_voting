package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// BodyCheckWrapper wraps the function with the BodyExtracted signature,
// checks for an empty and the limit of the body.
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

// ConvertBSToIS is a helper function to convert Byte Slice to Int Slice
func ConvertBSToIS(bSlice []byte) []int {
	var iSlice []int
	for _, b := range bSlice {
		iSlice = append(iSlice, int(b))
	}
	return iSlice
}

// ConvertISToBS is a helper function to convert Int Slice to Byte Slice
func ConvertISToBS(iSlice []int) []byte {
	var bSlice []byte
	for _, b := range iSlice {
		bSlice = append(bSlice, byte(b))
	}
	return bSlice
}

// RegexpStr matches the str with expr
func RegexpStr(expr string, str string) error {
	matched, err := regexp.MatchString(expr, str)

	if err != nil {
		return err
	}

	if matched != true {
		return errors.New("Match failed")
	}

	return nil

}

// IsIn checks if x is in [y]
func IsIn(x string, ys []string) bool {
	for _, y := range ys {
		if y == x {
			return true
		}
	}
	return false
}

// MethodWrapper wraps a handler func to respond only to the given method
func MethodWrapper(requestType string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.Method) != strings.ToUpper(requestType) {
			http.Error(w, "Method not allowed :-(", 400)
			return
		}
		fn(w, r)
	}
}

// FillBody reads the body from the req and fills it.
func (s *Service) FillBody() error {
	if s.Request.Body == nil {
		return errors.New("Empty Body")
	}

	body, err := ioutil.ReadAll(s.Request.Body)
	s.Request.Body.Close()
	if err != nil {
		return err
	}

	if len(body) > MaxReqBody {
		return errors.New("Body too long :-(")
	}

	s.Body = body
	return nil
}

func (s *Service) Error(str string, statusCode int) {
	s.Writer.Header().Set("Content-Type", "application/json")
	s.Writer.WriteHeader(statusCode)
	json.NewEncoder(s.Writer).Encode(BasicResponse{str, statusCode})
}
