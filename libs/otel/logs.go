package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func newLoggerProvider(cfg Config, ctx context.Context, res *resource.Resource) (*log.LoggerProvider, error) {
	exporter, err := stdoutlog.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP log exporter: %w", err)
	}

	processor := log.NewBatchProcessor(exporter)

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(processor),
		log.WithResource(res),
	)

	global.SetLoggerProvider(loggerProvider)

	return loggerProvider, nil
}
