package main

import (
	"context"
	"log"

	"github.com/erknas/wt-weapons/internal/api"
	"github.com/erknas/wt-weapons/internal/config"
	"github.com/erknas/wt-weapons/internal/logger"
	"github.com/erknas/wt-weapons/internal/storage"
)

func main() {
	var (
		ctx    = context.Background()
		cfg    = config.Load()
		logger = logger.New(cfg.Env)
	)

	storage, err := storage.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := storage.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	server := api.NewServer(logger, storage)
	if err := server.Start(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}
