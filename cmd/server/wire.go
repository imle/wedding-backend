//+build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/google/wire"

	"wedding/pkg/apiv1"
	"wedding/pkg/database"
	"wedding/pkg/health"
	"wedding/pkg/server"
)

func InitializeServer(_ context.Context) (*http.Server, func(), error) {
	wire.Build(
		ProvideRedisOptions,
		ProvideServerConfig,
		ProvideEntConfig,
		database.ProvideEntClient,
		health.NewMonitor,
		apiv1.NewRSVP,
		server.ProvideRedisSessionStore,
		server.ProvideEngine,
		server.ProvideHttpServer,
	)

	return &http.Server{}, func() {}, nil
}
