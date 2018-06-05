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
func BodyCheckWrapper(fn func(Service)) func(Service) {
	return func(s Service) {
		if s.Request.Body == nil {
			s.Tell("Please send a request body", 400)
			return
		}
		body, err := ioutil.ReadAll(s.Request.Body)
		s.Request.Body.Close()
		if err != nil {
			s.Tell(err.Error(), 400)
			return
		}

		if len(body) > MaxReqBody {
			s.Tell("Body too long :-(", 400)
			return
		}
		s.Body = body
		fn(s)
	}
}

// CreateService maps http.HandlerFunc -> Service
func CreateService(fn func(Service)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := Service{Writer: w, Request: r}
		fn(s)
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

// IsIn checks if x is in [y], x and y need to ASCII only
func IsIn(x string, y string) bool {
	chkMap := make(map[byte]bool)
	for i := 0; i < len(x); i++ {
		chkMap[x[i]] = false
	}
	for i := 0; i < len(y); i++ {
		if _, ok := chkMap[y[i]]; ok {
			return true
		}
	}
	return false
}

// MethodWrapper wraps a handler func to respond only to the given method
func MethodWrapper(requestType string, fn func(Service)) func(Service) {
	return func(s Service) {
		if strings.ToUpper(s.Request.Method) != strings.ToUpper(requestType) {
			s.Tell("Method not allowed :-(", 400)
			return
		}
		fn(s)
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

// Tell sets the status and "tells" the message
func (s *Service) Tell(str string, statusCode int) {
	s.Encode(BasicResponse{str, statusCode}, statusCode)
}

// Encode sets the status gives the result described
func (s *Service) Encode(data interface{}, statusCode int) {
	s.Writer.Header().Set("Content-Type", "application/json")
	s.Writer.WriteHeader(statusCode)
	json.NewEncoder(s.Writer).Encode(data)
}
