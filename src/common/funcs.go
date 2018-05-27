package common

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
	"strings"
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


func MethodWrapper(requestType string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.Method) != strings.ToUpper(requestType) {
			http.Error(w, "Method not allowed :-(", 400)
			return
		}
		fn(w, r)
	}
}




