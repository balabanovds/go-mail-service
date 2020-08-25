package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/balabanovds/mail-service/internal/mailer"
	"go.uber.org/zap"
)

func (s *ApiServer) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, map[string]bool{"alive": true})
	}
}

func (s *ApiServer) handleNewMail() http.HandlerFunc {

	type request struct {
		To      []string `json:"to"`
		Subject string   `json:"subject"`
		Body    string   `json:"body"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.serverError(w, err)
			return
		}

		err := s.mailer.Send(mailer.NewMail(req.To, req.Subject, req.Body))
		if err != nil {
			s.clientError(w, r, http.StatusNotAcceptable, err)
			return
		}

		zap.L().Info("mail has been sent",
			zapRequestID(r),
			zap.Strings("to", req.To),
			zap.String("subject", req.Subject),
		)
	}
}
