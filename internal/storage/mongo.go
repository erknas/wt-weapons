package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/erknas/wt-weapons/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	coll *mongo.Collection
}

func New(cfg *config.Config) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.AuthSource)

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %s", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping error: %s", err)
	}

	coll := client.Database(cfg.DBName).Collection(cfg.CollName)

	return &Storage{coll: coll}, nil
}
