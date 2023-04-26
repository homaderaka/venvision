package storage

import (
	"context"
	"venvision/protofiles"
)

type Storage interface {
	InsertOne(ctx context.Context, p *protofiles.Product) error
	FindByUUID(ctx context.Context, uuid string) (*protofiles.Product, error)
	FindAll(ctx context.Context) ([]*protofiles.Product, error)
	Update(ctx context.Context, uuid string, p *protofiles.Product) error
	Delete(ctx context.Context, uuid string) error
	DeleteAll(ctx context.Context) error
}
