package service

import (
	"context"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type ClientService struct {
	repo domain.ClientRepository
}

func NewClientService(repo domain.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Create(ctx context.Context, c *domain.Client) error {
	return s.repo.Create(ctx, c)
}

func (s *ClientService) List(ctx context.Context) ([]*domain.Client, error) {
	return s.repo.List(ctx)
}
