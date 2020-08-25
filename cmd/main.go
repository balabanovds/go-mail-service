package main

import (
	"flag"
	"log"

	"github.com/balabanovds/mail-service/internal/apiserver"
	"github.com/balabanovds/mail-service/internal/mailer"
	"go.uber.org/zap"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "cfg", "config/default.yml", "config file")
}

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync() //nolint:errcheck

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	cfg, err := LoadConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	server := apiserver.NewApiServer(cfg.Server, mailer.NewMailer(cfg.Mail))
	if err := server.Start(); err != nil {
		zap.L().Error("failed to start server")
	}
}
