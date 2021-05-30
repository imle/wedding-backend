package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	IsHealthy(ctx context.Context) bool
}

type Monitor struct {
	services map[string]Service
}

func NewMonitor() *Monitor {
	return &Monitor{
		services: map[string]Service{},
	}
}

func (h *Monitor) RegisterService(name string, ha Service) {
	h.services[name] = ha
}

func (h *Monitor) IsHealthy(ctx context.Context) (map[string]bool, bool) {
	overallHealthy := true
	sh := map[string]bool{}

	done := make(chan struct{})
	defer close(done)

	go func() {
		select {
		case <-ctx.Done():
			if len(h.services) != len(sh) {
				overallHealthy = false
			}
		case <-done:
		}
	}()

	for name, ha := range h.services {
		sh[name] = ha.IsHealthy(ctx)
		if !sh[name] && overallHealthy {
			overallHealthy = false
		}
	}
	return sh, overallHealthy
}

func (h *Monitor) Register(g *gin.RouterGroup) {
	g.GET("", func(c *gin.Context) {
		healthMap, overallHealthy := h.IsHealthy(c.Request.Context())

		status := http.StatusOK
		if !overallHealthy {
			status = http.StatusServiceUnavailable
		}

		c.JSON(status, healthMap)
	})
}
