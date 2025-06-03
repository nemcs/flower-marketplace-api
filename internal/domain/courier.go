package domain

import "context"

type Courier struct {
    ID        int64  `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
}

type CourierRepository interface {
    Create(ctx context.Context, c *Courier) error
    List(ctx context.Context) ([]*Courier, error)
}
