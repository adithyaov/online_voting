package common

import (
	"net/http"
)

var MaxReqBody = 4000

type BodyExtracted func (http.ResponseWriter, *http.Request, *[]byte)
