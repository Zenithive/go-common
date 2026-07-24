package logx

import (
	"log/slog"
	"os"
)

func New(opts ...Option) *slog.Logger {
	cfg := cfg{
		Level: slog.LevelInfo,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	h := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: cfg.Level,
		},
	)

	log := slog.New(h)

	if cfg.Service != "" {
		log = log.With(
			slog.String("service", cfg.Service),
		)
	}

	if cfg.Env != "" {
		log = log.With(
			slog.String("env", cfg.Env),
		)
	}

	return log
}
