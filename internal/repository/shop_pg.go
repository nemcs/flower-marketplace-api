// Реализация Postgres-репозитория
package repository

import (
	"context"
	"database/sql"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type ShopRepoPostgres struct {
	db *sql.DB
}

func NewShopRepoPostgres(db *sql.DB) *ShopRepoPostgres {
	return &ShopRepoPostgres{db: db}
}

func (r *ShopRepoPostgres) Create(ctx context.Context, shop *domain.Shop) error {
	query := `INSERT INTO shops (name, address) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRowContext(ctx, query, shop.Name, shop.Address).Scan(&shop.ID)
}

func (r *ShopRepoPostgres) List(ctx context.Context) ([]*domain.Shop, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, address FROM shops`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []*domain.Shop
	for rows.Next() {
		var s domain.Shop
		if err := rows.Scan(&s.ID, &s.Name, &s.Address); err != nil {
			return nil, err
		}
		shops = append(shops, &s)
	}

	return shops, nil
}
