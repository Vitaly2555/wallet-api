package main

import (
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"

	"wallet/internal/handler"
	"wallet/internal/repository"
	"wallet/internal/usecase"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func main() {
	log := zerolog.New(os.Stderr)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to connect postgres")
	}
	repo := repository.NewWalletPostgres(db)
	uc := usecase.NewWalletUsecase(repo)
	h := handler.NewWalletHandler(uc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpChErr := httpStart(port, h.MakeMuxRouter())

	log.Info().Stack().Msg("wallet api started")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:
		log.Info().Msgf("shutting down by signal")
	case err = <-httpChErr:
		log.Panic().Stack().Err(err).Msg("http server")
	}
}

func httpStart(port string, h http.Handler) <-chan error {
	errCh := make(chan error)
	go func() {
		errCh <- errors.WithStack(http.ListenAndServe(":"+port, h))
		close(errCh)
	}()
	return errCh
}
