package repository

import (
	"context"
	"database/sql"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
)

type ClientRepoPostgres struct {
	db *sql.DB
}

func NewClientRepoPostgres(db *sql.DB) *ClientRepoPostgres {
	return &ClientRepoPostgres{db: db}
}

func (r *ClientRepoPostgres) Create(ctx context.Context, c *domain.Client) error {
	query := `INSERT INTO clients (name, phone) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRowContext(ctx, query, c.Name, c.Phone).Scan(&c.ID)
}

func (r *ClientRepoPostgres) List(ctx context.Context) ([]*domain.Client, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, phone FROM clients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []*domain.Client
	for rows.Next() {
		var c domain.Client
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
			return nil, err
		}
		clients = append(clients, &c)
	}

	return clients, nil
}
