package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoOpts struct {
	URL     string
	Timeout time.Duration
}

func NewMongo(ctx context.Context, opts MongoOpts) (*mongo.Client, error) {
	clientOpts := options.Client().
		ApplyURI(opts.URL)

	if opts.Timeout > 0 {
		clientOpts.ConnectTimeout = &opts.Timeout
	}

	client, err := mongo.Connect(
		ctx,
		clientOpts,
	)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
