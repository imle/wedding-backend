package main

import (
	"wedding/pkg/server"
)

func ProvideServerConfig() *server.Config {
	return &server.Config{
		Addr:          serverAddress,
		Environment:   environment,
		DevOrigin:     devOrigin.Value(),
		SessionSecret: redisSessionSecret,
	}
}
