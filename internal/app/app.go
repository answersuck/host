package app

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/exp/slog"

	"github.com/ysomad/answersuck/internal/config"
	"github.com/ysomad/answersuck/internal/pkg/httpserver"
	"github.com/ysomad/answersuck/internal/pkg/pgclient"
	playerpg "github.com/ysomad/answersuck/internal/postgres/player"
	tagpg "github.com/ysomad/answersuck/internal/postgres/tag"
	playersvc "github.com/ysomad/answersuck/internal/service/player"
	"github.com/ysomad/answersuck/internal/twirp"
	authtwirpv1 "github.com/ysomad/answersuck/internal/twirp/auth/v1"
	playertwirpv1 "github.com/ysomad/answersuck/internal/twirp/player/v1"
	tagtwirpv1 "github.com/ysomad/answersuck/internal/twirp/tag/v1"
)

func logFatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

func Run(conf *config.Config, flags Flags) { //nolint:funlen // main func
	if flags.Migrate {
		mustRunMigrations(conf.PG.URL)
	}

	pgClient, err := pgclient.New(
		conf.PG.URL,
		pgclient.WithMaxConns(conf.PG.MaxConns),
	)
	if err != nil {
		logFatal("pgclient.New", err)
	}

	// player
	playerPostgres := playerpg.NewRepository(pgClient)
	playerService := playersvc.NewService(playerPostgres)

	type playerUseCase struct {
		playerpg.Repository
		playersvc.Service
	}

	playerHandlerV1 := playertwirpv1.NewHandler(&playerUseCase{*playerPostgres, *playerService})

	// tag
	tagPostgres := tagpg.NewRepository(pgClient)
	tagHandlerV1 := tagtwirpv1.NewHandler(tagPostgres)

	// auth
	authHandlerV1 := authtwirpv1.NewHandler()

	// http
	mux := twirp.NewMux([]twirp.Handler{
		playerHandlerV1,
		tagHandlerV1,
		authHandlerV1,
	})

	srv := httpserver.New(mux, httpserver.WithPort(conf.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("received signal from httpserver", slog.String("signal", s.String()))
	case err := <-srv.Notify():
		slog.Info("got error from http server notify", slog.String("error", err.Error()))
	}

	if err := srv.Shutdown(); err != nil {
		slog.Info("got error on http server shutdown", slog.String("error", err.Error()))
	}
}
