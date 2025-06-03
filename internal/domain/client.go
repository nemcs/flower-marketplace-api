package domain

import "context"

type Client struct {
    ID    int64  `json:"id"`
    Name  string `json:"name"`
    Phone string `json:"phone"`
}

type ClientRepository interface {
    Create(ctx context.Context, c *Client) error
    List(ctx context.Context) ([]*Client, error)
}
