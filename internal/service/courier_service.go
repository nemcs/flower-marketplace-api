package service

import (
	"context"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type CourierService struct {
	repo domain.CourierRepository
}

func NewCourierService(repo domain.CourierRepository) *CourierService {
	return &CourierService{repo: repo}
}

func (s *CourierService) Create(ctx context.Context, c *domain.Courier) error {
	return s.repo.Create(ctx, c)
}

func (s *CourierService) List(ctx context.Context) ([]*domain.Courier, error) {
	return s.repo.List(ctx)
}
