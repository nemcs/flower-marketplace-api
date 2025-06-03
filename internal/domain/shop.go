// Entity + интерфейс репозитория
package domain

import "context"

type Shop struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ShopRepository interface {
	Create(ctx context.Context, shop *Shop) error
	List(ctx context.Context) ([]*Shop, error)
}
