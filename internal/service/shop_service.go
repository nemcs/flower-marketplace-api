//Бизнес-логика (use case)

package service

import (
	"context"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type ShopService struct {
	repo domain.ShopRepository
}

func NewShopService(repo domain.ShopRepository) *ShopService {
	return &ShopService{repo: repo}
}

func (s *ShopService) Create(ctx context.Context, shop *domain.Shop) error {
	return s.repo.Create(ctx, shop)
}

func (s *ShopService) List(ctx context.Context) ([]*domain.Shop, error) {
	return s.repo.List(ctx)
}
