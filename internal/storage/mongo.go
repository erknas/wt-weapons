package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/erknas/wt-weapons/internal/config"
	"github.com/erknas/wt-weapons/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const ctxTimeout = 5

type Storage struct {
	coll *mongo.Collection
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*ctxTimeout)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.AuthSource)

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	coll := client.Database(cfg.DBName).Collection(cfg.CollName)

	return &Storage{coll: coll}, nil
}

func (s *Storage) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	filter := bson.M{"category": category}

	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var weapons []*types.Weapon

	if err := cur.All(ctx, &weapons); err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return weapons, nil
}

func (s *Storage) AddWeapon(ctx context.Context, weapon *types.Weapon) (bson.ObjectID, error) {
	res, err := s.coll.InsertOne(ctx, weapon)
	if err != nil {
		return bson.NilObjectID, err
	}

	id := res.InsertedID

	return id.(bson.ObjectID), nil
}

func (s *Storage) UpdateWeapon(ctx context.Context, name string, weapon *types.Weapon) error {
	filter := bson.M{"name": name}
	update := bson.M{"$set": weapon}

	res, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("%s doesn't exist", name)
	}

	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*ctxTimeout)
	defer cancel()

	return s.coll.Database().Client().Disconnect(ctx)
}
