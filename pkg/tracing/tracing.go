package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const TraceID = "x-trace-id"

func InitJaeger(jaegerURL, serviceName string) (opentracing.Tracer, func(), error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerURL,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	log.Info().Msg("Jaeger connected with service: " + serviceName)

	return tracer, func() {
		if err = closer.Close(); err != nil {
			log.Error().Msgf("failed to close tracer, error: %v", err)
		}
	}, nil
}
