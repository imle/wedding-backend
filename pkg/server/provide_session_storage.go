package server

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	redissessions "github.com/gin-contrib/sessions/redis"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func ProvideRedisSessionStore(options *redis.Options, cfg *Config) (sessions.Store, error) {
	if cfg.SessionSecret == "" {
		return nil, fmt.Errorf("redis-session-secret cannot be empty")
	}

	db, err := redissessions.NewStoreWithDB(
		10,
		options.Network,
		options.Addr,
		options.Password,
		strconv.Itoa(options.DB),
		[]byte(cfg.SessionSecret),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error setting up redis session store")
	}

	return db, nil
}
