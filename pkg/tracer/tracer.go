package tracer

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.26.0"
	"log"
)

const (
	ProviderNameApiGateway  = "ApiGateway"
	ProviderNameAuthService = "AuthService"
	TraceId                 = "x-trace-id"
)

func Init(jaegerURL string, serviceName string) (*trace.TracerProvider, func(context.Context), error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		return nil, nil, fmt.Errorf("initialize exporter: %w", err)
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize resource: %w", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(r),
	)

	otel.SetTracerProvider(tp)

	return tp, func(ctx context.Context) {
		if err = tp.Shutdown(ctx); err != nil {
			log.Fatalf("Ошибка при завершении Tracer Provider: %v", err)
		}
	}, nil
}
