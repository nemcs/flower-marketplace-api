package domain

import "context"

type Order struct {
    ID        int64  `json:"id"`
    ClientID  int64  `json:"client_id"`
    ShopID    int64  `json:"shop_id"`
    CourierID int64  `json:"courier_id"`
    Address   string `json:"address"`
    Status    string `json:"status"` // e.g., "pending", "delivered"
}

type OrderRepository interface {
    Create(ctx context.Context, o *Order) error
    List(ctx context.Context) ([]*Order, error)
}
