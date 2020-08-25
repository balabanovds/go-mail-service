package main

import (
	"context"
	"time"

	"github.com/balabanovds/mail-service/internal/apiserver"
	"github.com/balabanovds/mail-service/internal/mailer"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	Mail   *mailer.Config    `config:"mail"`
	Server *apiserver.Config `config:"server"`
}

func LoadConfig(filename string) (*Config, error) {
	loader := confita.NewLoader(
		env.NewBackend(),
		file.NewBackend(filename),
		flags.NewBackend(),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cfg Config
	err := loader.Load(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
