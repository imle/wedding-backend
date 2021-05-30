package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"wedding/pkg/apiv1"
	"wedding/pkg/health"
	"wedding/pkg/wedding"
)

func ProvideEngine(
	cfg *Config,
	hm *health.Monitor,
	rsvpApiV1 *apiv1.RSVP,
) (*gin.Engine, error) {
	switch cfg.Environment {
	case wedding.EnvironmentProduction, wedding.EnvironmentStaging:
		gin.SetMode(gin.ReleaseMode)
	case wedding.EnvironmentTest:
		gin.SetMode(gin.TestMode)
	case wedding.EnvironmentDevelopment, wedding.EnvironmentLocal:
		gin.SetMode(gin.DebugMode)
	default:
		log.Fatal("invalid environment")
	}

	engine := gin.New()
	engine.RedirectTrailingSlash = false
	engine.RemoveExtraSlash = true

	// panic recovery middleware
	engine.Use(gin.Recovery())

	// Create internal service monitoring router.
	// We do this before registering the logging and tracing to avoid excessive logs on these endpoints.
	monitoringRouter := engine.Group("/_")
	hm.Register(monitoringRouter.Group("/health"))

	// otel middleware
	engine.Use(otelgin.Middleware("wedding"))
	// logrus middleware
	engine.Use(ginlogrus.Logger(log.StandardLogger()))

	// Setup CORS.
	if gin.Mode() != gin.ReleaseMode {
		config := cors.DefaultConfig()
		config.AllowCredentials = true
		config.AllowOrigins = append(config.AllowOrigins, cfg.DevOrigin...)
		engine.Use(cors.New(config))
	}

	// Create base router.
	router := engine.Group("/api/v1")

	// Register authed v1 route handlers.
	rsvpApiV1.Register(router.Group("/invitee"), router.Group("/invitee"))

	return engine, nil
}
