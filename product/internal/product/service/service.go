package service

import (
	"context"
	"venvision/pkg/logging"
	"venvision/protofiles"
)

type productService struct {
	logger logging.Logger
}

func NewProductService(logger *logging.Logger) (*productService, error) {
	return &productService{
		logger: *logger,
	}, nil
}

func (s *productService) GetAll(ctx context.Context, req *protofiles.GetAllRequest) (*protofiles.GetAllResponse, error) {
	return &protofiles.GetAllResponse{}, nil
}

func (s *productService) GetByUUID(ctx context.Context, req *protofiles.GetByUUIDRequest) (*protofiles.Product, error) {
	// Implement the logic to get a product by UUID here.
	return &protofiles.Product{}, nil
}
