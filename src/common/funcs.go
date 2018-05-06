package common

import (
	"net/http"
)


func BodyCheckWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		fn(w, r)
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






