package api

import (
	"net/http"
)

type API interface {
	Start()

	CreateURLHandler(w http.ResponseWriter, r *http.Request)

	RedirectHandler(w http.ResponseWriter, r *http.Request)
}
