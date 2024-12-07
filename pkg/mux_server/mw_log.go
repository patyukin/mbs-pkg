package mux_server

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestID, err := uuid.NewUUID()
			if err != nil {
				log.Info().Msgf("unable to generate request ID in LoggingMiddleware: %v", err)
				http.Error(w, "Unable to generate request ID", http.StatusInternalServerError)
				return
			}

			log.Info().Msgf(
				"Request ID: %s, Path: %s, Query Params: %s", requestID, r.URL.Path, r.URL.RawQuery,
			)

			startTime := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(startTime)

			log.Info().Msgf("Completed Request ID: %s in %v", requestID, duration)
		},
	)
}
