package domain

import (
	"context"

	"gorm.io/gorm"
)

type Order struct {
	ID         int64   `json:"id"`
	ProductID  int64   `json:"product_id"`
	CreatorID  int64   `json:"creator_id"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	Status     bool    `json:"status"`
	CreateTime int64   `json:"create_time"`
	UpdateTime int64   `json:"update_time"`
}
type RepositoryOrder interface {
	CreateOrder(ctx context.Context, tx *gorm.DB, order *Order) error
	GetOrdersByUserId(ctx context.Context, creator_id int64) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error
	DeleteOrder(ctx context.Context, id int64) error
}
