package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func Meter(name string) metric.Meter {
	return otel.Meter(name)
}

func Counter(meter metric.Meter, name string) (metric.Int64Counter, error) {
	return meter.Int64Counter(name)
}

func Add(ctx context.Context, counter metric.Int64Counter, value int64) {
	counter.Add(ctx, value)
}
