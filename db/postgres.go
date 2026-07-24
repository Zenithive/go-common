package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGOpts struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Name     string
	SSLMode  string

	MaxOpen     int32
	MaxIdle     int32
	MaxIdleTime time.Duration
}

func NewPostgres(ctx context.Context, opts PGOpts) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}

	cfg.ConnConfig.Host = opts.Host
	cfg.ConnConfig.Port = opts.Port
	cfg.ConnConfig.User = opts.User
	cfg.ConnConfig.Password = opts.Password
	cfg.ConnConfig.Database = opts.Name

	if opts.SSLMode != "" {
		cfg.ConnConfig.RuntimeParams["sslmode"] = opts.SSLMode
	}

	if opts.MaxOpen > 0 {
		cfg.MaxConns = opts.MaxOpen
	}

	if opts.MaxIdle > 0 {
		cfg.MinConns = opts.MaxIdle
	}

	if opts.MaxIdleTime > 0 {
		cfg.MaxConnIdleTime = opts.MaxIdleTime
	}

	return pgxpool.NewWithConfig(ctx, cfg)
}
