package common

import (
	"net/http"
)

type BasicResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
}

type BodyExtracted func (http.ResponseWriter, *http.Request, *[]byte)

