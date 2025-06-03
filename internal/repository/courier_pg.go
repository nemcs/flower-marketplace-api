package repository

import (
	"context"
	"database/sql"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type CourierRepoPostgres struct {
	db *sql.DB
}

func NewCourierRepoPostgres(db *sql.DB) *CourierRepoPostgres {
	return &CourierRepoPostgres{db: db}
}

func (r *CourierRepoPostgres) Create(ctx context.Context, c *domain.Courier) error {
	query := `INSERT INTO couriers (first_name, last_name, phone) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRowContext(ctx, query, c.FirstName, c.LastName, c.Phone).Scan(&c.ID)
}

func (r *CourierRepoPostgres) List(ctx context.Context) ([]*domain.Courier, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, first_name, last_name, phone FROM couriers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var couriers []*domain.Courier
	for rows.Next() {
		var c domain.Courier
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone); err != nil {
			return nil, err
		}
		couriers = append(couriers, &c)
	}

	return couriers, nil
}
