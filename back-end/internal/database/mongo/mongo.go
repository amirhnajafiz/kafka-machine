package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewConnection creates a connection to our mongodb
func NewConnection(cfg Config) (*mongo.Database, error) {
	ctx := context.Background()

	// create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URL))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	// ping mongodb
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	return client.Database(cfg.Database), nil
}
