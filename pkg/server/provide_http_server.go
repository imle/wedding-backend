package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"wedding/pkg/wedding"
)

type Config struct {
	Addr          string
	DevOrigin     []string
	Environment   wedding.Environment
	SessionSecret string
}

func ProvideHttpServer(engine *gin.Engine, cfg *Config) *http.Server {
	return &http.Server{
		Addr:    cfg.Addr,
		Handler: engine,
	}
}
