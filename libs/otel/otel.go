package otel

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Telemetry struct {
	lp *log.LoggerProvider
	mp *metric.MeterProvider
	tp *trace.TracerProvider
}

func NewTelemetry(ctx context.Context) (*Telemetry, error) {
	cfg, err := newConfigFromEnv()
	if err != nil {
		fmt.Println("failed to load telemetry config")
		os.Exit(1)
	}

	rp := newResource(cfg.ServiceName, cfg.ServiceVersion)

	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up logger provider.
	loggerProvider, err := newLoggerProvider(cfg, ctx, rp)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	// Set up meter provider.
	meterProvider, err := newMeterProvider(cfg, ctx, rp)
	if err != nil {
		return nil, fmt.Errorf("failed to create meter: %w", err)
	}

	// Set up trace provider.
	tracerProvider, err := newTracerProvider(cfg, ctx, rp)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer: %w", err)
	}

	return &Telemetry{
		lp: loggerProvider,
		mp: meterProvider,
		tp: tracerProvider,
	}, nil
}

// Shutdown shuts down the logger, meter, and tracer.
func (t *Telemetry) Shutdown(ctx context.Context) {
	t.lp.Shutdown(ctx)
	t.mp.Shutdown(ctx)
	t.tp.Shutdown(ctx)
}

func Middlewares() gin.HandlerFunc {
	cfg, err := newConfigFromEnv()
	if err != nil {
		fmt.Println("failed to load telemetry config")
		os.Exit(1)
	}

	println("aaa", cfg.ServiceName)

	return otelgin.Middleware(cfg.ServiceName)
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newResource(serviceName string, serviceVersion string) *resource.Resource {
	hostName, _ := os.Hostname()

	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion(serviceVersion),
		semconv.HostName(hostName),
	)
}
