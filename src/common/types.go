package common

import (
	"net/http"
	"sync"
)

// BasicResponse describes the basic response with a secondary status code.
type BasicResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// BodyExtracted is a type with writer, req and body (in bytes).
type BodyExtracted func(http.ResponseWriter, *http.Request, *[]byte)

// ThreadSafeType is any data with mutex mutex type
type ThreadSafeType struct {
	Mutex sync.Mutex
	Data  interface{}
}
