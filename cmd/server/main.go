package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"wedding/internal/pkg/opentelemetry"
	"wedding/pkg/wedding"
)

var (
	pgHost     string
	pgPort     int
	pgUsername string
	pgDatabase string
	pgPassword string
	pgSSLMode  string

	serverAddress string
	logLevel      string
	environment   wedding.Environment
	devOrigin     cli.StringSlice

	redisUrl           string
	redisDb            int
	redisSessionSecret string

	otelCollectorInsecure bool
	otelCollectorEndpoint string
	otelCollectorPeriod   time.Duration
)

func main() {
	app := cli.App{
		Name:      "wedding",
		ArgsUsage: "[[hostname]:port]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Destination: &serverAddress,
				EnvVars:     []string{"WEDDING_ADDRESS"},
				Usage:       "[hostname]:port",
				Value:       ":8080",
			},
			&cli.StringFlag{
				Name:        "log-level",
				Destination: &logLevel,
				EnvVars:     []string{"WEDDING_LOG_LEVEL"},
				Value:       log.InfoLevel.String(),
			},
			&cli.StringFlag{
				Name:        "environment",
				Destination: (*string)(&environment),
				EnvVars:     []string{"WEDDING_ENVIRONMENT"},
				Value:       string(wedding.EnvironmentProduction),
			},
			&cli.StringSliceFlag{
				Name:        "development-origin",
				Destination: &devOrigin,
				EnvVars:     []string{"WEDDING_DEV_ORIGIN"},
				Value:       cli.NewStringSlice("http://localhost:3000"),
			},
			&cli.StringFlag{
				Name:        "pg-host",
				Destination: &pgHost,
				EnvVars:     []string{"WEDDING_PG_HOST"},
				Value:       "localhost",
			},
			&cli.IntFlag{
				Name:        "pg-port",
				Destination: &pgPort,
				EnvVars:     []string{"WEDDING_PG_PORT"},
				Value:       5432,
			},
			&cli.StringFlag{
				Name:        "pg-username",
				Destination: &pgUsername,
				EnvVars:     []string{"WEDDING_PG_USERNAME"},
				Value:       "wedding",
			},
			&cli.StringFlag{
				Name:        "pg-database",
				Destination: &pgDatabase,
				EnvVars:     []string{"WEDDING_PG_DATABASE"},
				Value:       "wedding",
			},
			&cli.StringFlag{
				Name:        "pg-password",
				Destination: &pgPassword,
				EnvVars:     []string{"WEDDING_PG_PASSWORD"},
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "pg-sslmode",
				Destination: &pgSSLMode,
				EnvVars:     []string{"WEDDING_PG_SSL_MODE"},
				Value:       "disable",
			},
			&cli.StringFlag{
				Name:        "redis-url",
				Destination: &redisUrl,
				EnvVars:     []string{"WEDDING_REDIS_URL"},
				Value:       "localhost:6379",
			},
			&cli.IntFlag{
				Name:        "redis-db",
				Destination: &redisDb,
				EnvVars:     []string{"WEDDING_REDIS_DB"},
				Value:       0,
			},
			&cli.StringFlag{
				Name:        "redis-session-secret",
				Destination: &redisSessionSecret,
				EnvVars:     []string{"WEDDING_REDIS_SESSION_SECRET"},
				Value:       "",
			},
			&cli.BoolFlag{
				Name:        "otel-collector-insecure",
				Destination: &otelCollectorInsecure,
				EnvVars:     []string{"WEDDING_OTEL_COLLECTOR_INSECURE"},
				Value:       false,
			},
			&cli.StringFlag{
				Name:        "otel-collector-endpoint",
				Destination: &otelCollectorEndpoint,
				EnvVars:     []string{"WEDDING_OTEL_COLLECTOR_ENDPOINT"},
				Value:       "localhost:30080",
			},
			&cli.DurationFlag{
				Name:        "otel-collector-period",
				Destination: &otelCollectorPeriod,
				EnvVars:     []string{"WEDDING_OTEL_COLLECTOR_PERIOD"},
				Value:       2 * time.Second,
			},
		},
		Commands: []*cli.Command{
			&ImportGuestList,
			&GenerateMigrations,
		},
		Action: func(c *cli.Context) error {
			level, err := log.ParseLevel(logLevel)
			if err != nil {
				return err
			}
			log.SetLevel(level)
			log.SetReportCaller(true)
			log.AddHook(&opentelemetry.LogrusTraceHook{})
			if !environment.Is(wedding.EnvironmentProduction, wedding.EnvironmentStaging) {
				log.SetFormatter(&log.JSONFormatter{})
			}

			stopOtelProvider, err := opentelemetry.InitOtelProvider(c.Context, opentelemetry.PFOtelProviderConfig{
				ServiceName:           c.App.Name,
				OtelCollectorEndpoint: otelCollectorEndpoint,
				OtelCollectorInsecure: otelCollectorInsecure,
				OtelCollectorPeriod:   otelCollectorPeriod,
			})
			if err != nil {
				return err
			}

			srv, closer, err := InitializeServer(c.Context)
			if err != nil {
				return err
			}
			defer closer()

			// Initializing the server in a goroutine to enable graceful handling of shutdown.
			go func() {
				log.Infof("server listening on %s", srv.Addr)
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatal("error while serving: ", err)
				}
			}()

			<-c.Context.Done()
			log.Info("shutting down server...")

			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err = srv.Shutdown(shutdownCtx)
			if err != nil {
				return err
			}

			log.Info("server exited")

			stopOtelProvider(shutdownCtx)

			log.Info("tracing stopped")

			return nil
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGQUIT)
	go func() {
		<-quit
		cancel()
	}()

	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
