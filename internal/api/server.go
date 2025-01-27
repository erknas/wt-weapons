package api

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-weapons/internal/config"
	"github.com/erknas/wt-weapons/internal/logger"
	"github.com/erknas/wt-weapons/internal/storage"
	"github.com/erknas/wt-weapons/lib"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const ctxTimeout = 10

type Server struct {
	log    *slog.Logger
	storer storage.Storer
}

func NewServer(log *slog.Logger, storer storage.Storer) *Server {
	return &Server{
		log:    log,
		storer: storer,
	}
}

func (s *Server) Start(ctx context.Context, cfg *config.Config) error {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(logger.NewMwLogger(s.log))

	s.registerRoutes(router)

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	quitch := make(chan os.Signal, 1)
	signal.Notify(quitch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	s.log.Info("starting server", "port", srv.Addr)

	<-quitch

	shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*ctxTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	s.log.Info("server shutdown")

	return nil
}

func (s *Server) registerRoutes(router *chi.Mux) {
	router.Get("/weapons/{category}", lib.MakeHTTPFunc(s.handleGetWeaponsByCategory))

	router.Route("/weapons", func(r chi.Router) {
		r.Post("/", lib.MakeHTTPFunc(s.handleAddWeapon))
		r.Put("/{name}", lib.MakeHTTPFunc(s.handleUpdateWeapon))
	})
}
