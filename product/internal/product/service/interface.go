package service

import (
	"context"
	"venvision/protofiles"
)

type Service interface {
	GetAll(ctx context.Context, req *protofiles.GetAllRequest) (*protofiles.GetAllResponse, error)
	GetByUUID(ctx context.Context, req *protofiles.GetByUUIDRequest) (*protofiles.Product, error)
}
