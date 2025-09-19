package otel

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServiceName          string `env:"SERVICE_NAME"      envDefault:""`
	ServiceVersion       string `env:"SERVICE_VERSION"   envDefault:"0.1.0"`
	OTELExporterEndpoint string `env:"OTEL_EXPORTER_ENDPOINT"   envDefault:"127.0.0.1:4317"`
	EnabledOTELLogger    bool   `env:"ENABLE_OTEL_LOGGER" envDefault:"false"`
	EnabledOTELTracer    bool   `env:"ENABLE_OTEL_TRACER" envDefault:"false"`
	EnabledOTELMeter     bool   `env:"ENABLE_OTEL_METER" envDefault:"false"`
}

func newConfigFromEnv() (Config, error) {
	otelConfig := Config{}
	if err := env.Parse(&otelConfig); err != nil {
		return Config{}, fmt.Errorf("failed to parse telemetry config: %w", err)
	}

	return otelConfig, nil
}
