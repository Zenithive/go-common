package logx

import "log/slog"

type Option func(*cfg)

type cfg struct {
	Service string
	Env     string
	Level   slog.Level
}

func WithService(name string) Option {
	return func(c *cfg) {
		c.Service = name
	}
}

func WithEnv(env string) Option {
	return func(c *cfg) {
		c.Env = env
	}
}

func WithLevel(level slog.Level) Option {
	return func(c *cfg) {
		c.Level = level
	}
}
