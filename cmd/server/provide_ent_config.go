package main

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"wedding/pkg/database"
)

func ProvideEntConfig() (*database.EntConfig, error) {
	connData := map[string]string{
		"host":     pgHost,
		"port":     strconv.Itoa(pgPort),
		"user":     pgUsername,
		"dbname":   pgDatabase,
		"password": pgPassword,
		"sslmode":  pgSSLMode,
	}
	var connectionString string
	for key, value := range connData {
		if value == "" {
			continue
		}
		connectionString += key + "=" + value + " "
	}

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	return &database.EntConfig{
		ConnectionString: connectionString,
		Debug:            level == log.DebugLevel,
	}, nil
}
