package apiserver

import (
	"fmt"
	"net/http"

	"github.com/balabanovds/mail-service/internal/mailer"
	"go.uber.org/zap"
)

type ApiServer struct {
	config *Config
	mailer mailer.MailSender
}

func NewApiServer(config *Config, mailer mailer.MailSender) *ApiServer {
	return &ApiServer{
		config,
		mailer,
	}
}

func (s *ApiServer) Start() error {
	zap.L().Info("starting server",
		zap.String("host", s.config.Host),
		zap.Int("port", s.config.Port),
	)

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	server := http.Server{
		Addr:     addr,
		Handler:  s.routes(),
		ErrorLog: zap.NewStdLog(zap.L()),
	}

	return server.ListenAndServe()
}
