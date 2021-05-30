package main

import (
	log "github.com/sirupsen/logrus"

	"wedding/pkg/database"
)

func ProvideEntConfig() (*database.EntConfig, error) {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	return &database.EntConfig{
		ConnectionString: dbConStr,
		Debug:            level == log.DebugLevel,
	}, nil
}
