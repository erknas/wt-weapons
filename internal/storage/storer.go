package storage

import (
	"context"

	"github.com/erknas/wt-weapons/types"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Storer interface {
	GetWeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
	AddWeapon(context.Context, *types.Weapon) (bson.ObjectID, error)
	UpdateWeapon(context.Context, string, *types.Weapon) error
}
