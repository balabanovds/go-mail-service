package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/balabanovds/mail-service/internal/mailer"
	"github.com/stretchr/testify/require"
)

var (
	cfg = Config{
		Host: "localhost",
		Port: 9001,
	}
)

func TestHealthCheckHandler(t *testing.T) {
	s := NewApiServer(&cfg, nil)
	req, err := http.NewRequest("GET", "/health-check", nil)
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	s.handleHealthCheck().ServeHTTP(rec, req)
	require.Equal(t, `{"alive":true}`, strings.TrimSpace(rec.Body.String()))
}

func TestNewMailHandler(t *testing.T) {
	type body struct {
		To      []string `json:"to"`
		Subject string   `json:"subject"`
		Body    string   `json:"body"`
	}

	goodBody := body{}
	// goodBody := body{Subject: "test", Body: "test"}
	bodyBytes, err := json.Marshal(goodBody)
	require.NoError(t, err)

	tests := []struct {
		name       string
		mailer     mailer.MailSender
		body       []byte
		statusCode int
	}{
		{
			name:       "error parsing request body",
			mailer:     nil,
			body:       []byte(""),
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "error returned by mailer",
			mailer:     &mailer.BadMailerStub{},
			body:       bodyBytes,
			statusCode: http.StatusNotAcceptable,
		},
		{
			name:       "no error returned from mailer",
			mailer:     &mailer.GoodMailerStub{},
			body:       bodyBytes,
			statusCode: http.StatusOK,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			s := NewApiServer(&cfg, tst.mailer)
			rec := httptest.NewRecorder()

			req, err := http.NewRequest("POST", "/new", bytes.NewReader(tst.body))
			require.NoError(t, err)

			s.handleNewMail().ServeHTTP(rec, req)
			require.Equal(t, tst.statusCode, rec.Result().StatusCode)
		})
	}
}
