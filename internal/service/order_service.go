package service

import (
	"context"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type OrderService struct {
	repo domain.OrderRepository
}

func NewOrderService(repo domain.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(ctx context.Context, o *domain.Order) error {
	return s.repo.Create(ctx, o)
}

func (s *OrderService) List(ctx context.Context) ([]*domain.Order, error) {
	return s.repo.List(ctx)
}
