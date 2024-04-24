package handlers

import "net/http"

type Routes struct {
	URL               string
	Method            string
	Function          func(w http.ResponseWriter, r *http.Request) error
	ReqAuthentication bool
}

