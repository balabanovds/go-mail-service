package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiServer) routes() http.Handler {
	r := mux.NewRouter()

	r.Use(jsonContent)
	r.Use(s.withRequestID)
	r.Use(logRequest)

	r.HandleFunc("/health-check", s.handleHealthCheck())
	r.HandleFunc("/new", s.handleNewMail()).Methods("POST")

	return r
}
