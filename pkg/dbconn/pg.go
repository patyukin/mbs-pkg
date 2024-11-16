package dbconn

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

func New(ctx context.Context, dsn string) (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	dbConn.SetConnMaxLifetime(time.Hour)

	log.Info().Msg("connected to db")

	err = dbConn.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	log.Info().Msg("pinged postgresql db")

	return dbConn, nil
}
