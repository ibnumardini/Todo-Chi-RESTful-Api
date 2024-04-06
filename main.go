package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"todo-chi-rest-api/modules/todo"
	"todo-chi-rest-api/pkg/config"
	"todo-chi-rest-api/pkg/database"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func init() {
	config.InitConfig()
}

func main() {
	ctx := context.Background()

	dbConn, err := database.OpenDbConnection()
	if err != nil {
		panic(err)
	}

	todoModule := todo.LoadModule(dbConn)

	r := chi.NewRouter()

	r.Mount("/api", todoModule)

	server := http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", config.C.Server.AppHost, config.C.Server.AppPort),
	}

	log.Info().Msg("starting up server...")

	go func() {
		log.Info().Msgf("server will run at: %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Err(err).Msg("failed init http server")
			panic(err)
		}
	}()

	// gracefully shutdown
	wait := gracefullyShutdown(ctx, time.Duration(5*time.Second), map[string]cleanUpFunc{
		"db": func(ctx context.Context) error {
			return dbConn.Close()
		},
		"server": func(ctx context.Context) error {
			return server.Close()
		},
	})
	<-wait

	log.Info().Msg("server Stopped")
}

type cleanUpFunc func(ctx context.Context) error

func gracefullyShutdown(ctx context.Context, timeout time.Duration, funcs map[string]cleanUpFunc) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s
		log.Info().Msg("shutting down")

		// set timeout to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Warn().Msgf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		// exec the cleanup funcs asynchronously to save time
		var wg sync.WaitGroup
		for key, val := range funcs {
			wg.Add(1)

			go func(key string, cleanFunc cleanUpFunc) {
				defer wg.Done()

				log.Info().Msgf("cleaning up: %s", key)
				if err := cleanFunc(ctx); err != nil {
					log.Warn().Msgf("%s: clean up failed: %s", key, err.Error())
					return
				}

				log.Info().Msgf("%s was shutdown gracefully", key)
			}(key, val)
		}
		wg.Wait()

		close(wait)
	}()

	return wait
}
