package opentelemetry

import (
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type LogrusTraceHook struct{}

func (l LogrusTraceHook) Levels() []log.Level {
	return log.AllLevels
}

func (l LogrusTraceHook) Fire(entry *log.Entry) error {
	if entry.Context == nil {
		return nil
	}

	span := trace.SpanFromContext(entry.Context)
	if span == nil || !span.SpanContext().TraceID().IsValid() {
		return nil
	}

	entry.Data["trace"] = span.SpanContext().TraceID().String()

	return nil
}
