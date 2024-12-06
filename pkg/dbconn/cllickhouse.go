package dbconn

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/rs/zerolog/log"
)

func NewClickhouse(ctx context.Context, dsn string) (*sql.DB, error) {
	dbConn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	dbConn.SetConnMaxLifetime(time.Hour)

	log.Info().Msg("connected to db")

	err = dbConn.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	log.Info().Msg("pinged clickhouse db")

	return dbConn, nil
}
