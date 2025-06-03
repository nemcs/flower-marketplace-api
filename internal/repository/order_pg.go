package repository

import (
	"context"
	"database/sql"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type OrderRepoPostgres struct {
	db *sql.DB
}

func NewOrderRepoPostgres(db *sql.DB) *OrderRepoPostgres {
	return &OrderRepoPostgres{db: db}
}

func (r *OrderRepoPostgres) Create(ctx context.Context, o *domain.Order) error {
	query := `INSERT INTO orders (client_id, shop_id, courier_id, address, status)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowContext(ctx, query, o.ClientID, o.ShopID, o.CourierID, o.Address, o.Status).Scan(&o.ID)
}

func (r *OrderRepoPostgres) List(ctx context.Context) ([]*domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, client_id, shop_id, courier_id, address, status FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.ClientID, &o.ShopID, &o.CourierID, &o.Address, &o.Status); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}

	return orders, nil
}
