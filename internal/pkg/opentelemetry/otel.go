package opentelemetry

import (
	"context"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

type ErrorHandler struct {
}

func (e *ErrorHandler) Handle(err error) {
	log.Error(errors.Wrap(err, "otel provider error"))
}

type PFOtelProviderConfig struct {
	ServiceName           string
	OtelCollectorEndpoint string
	OtelCollectorInsecure bool
	OtelCollectorPeriod   time.Duration
}

// InitOtelProvider creates a new trace provider instance and registers it as global trace provider.
func InitOtelProvider(ctx context.Context, config PFOtelProviderConfig) (func(context.Context), error) {
	if config.OtelCollectorEndpoint == "" {
		return nil, nil
	}

	// Custom error log handler.
	otel.SetErrorHandler(&ErrorHandler{})

	// Setup options for otel push collector.
	options := []otlpgrpc.Option{
		otlpgrpc.WithEndpoint(config.OtelCollectorEndpoint),
	}
	if config.OtelCollectorInsecure {
		options = append(options, otlpgrpc.WithInsecure())
	}

	// Create driver with above options.
	driver := otlpgrpc.NewDriver(options...)

	// Create exporter using the new driver.
	exp, err := otlp.NewExporter(ctx, driver)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create exporter")
	}

	// Set service name for exporter.
	nameRes := resource.NewWithAttributes(
		semconv.ServiceNameKey.String(config.ServiceName),
	)

	// BEGIN TRACING
	// Create span pre-processor and trace provider.
	bsp := trace.NewBatchSpanProcessor(exp)
	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(nameRes),
		trace.WithSpanProcessor(bsp),
	)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	// Set global trace provider.
	otel.SetTracerProvider(tracerProvider)
	// END TRACING

	// BEGIN METRICS
	// Create metrics processor.
	metricsProcessor := processor.New(
		simple.NewWithHistogramDistribution(
			histogram.WithExplicitBoundaries([]float64{
				0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000,
			}),
		),
		exp,
	)

	// Create metrics controller.
	cont := controller.New(
		metricsProcessor,
		controller.WithResource(nameRes),
		controller.WithExporter(exp),
		controller.WithCollectPeriod(2*time.Second),
	)

	// Start metrics controller.
	err = cont.Start(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not start metric controller")
	}
	// Set the global metrics provider.
	global.SetMeterProvider(cont.MeterProvider())
	// END METRICS

	return func(ctx context.Context) {
		err = cont.Stop(ctx)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to shutdown controller"))
		}

		// Shutdown will flush any remaining spans and shut down the exporter.
		err = tracerProvider.Shutdown(ctx)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to shutdown TracerProvider"))
		}
	}, nil
}
