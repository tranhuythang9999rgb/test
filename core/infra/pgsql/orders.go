package pgsql

import (
	"ap_sell_products/core/domain"
	"ap_sell_products/core/infra"
	"context"

	"gorm.io/gorm"
)

type CollectionOrder struct {
	collection *gorm.DB
}

func NewCollectionOrder(pd *infra.PostGresql) domain.RepositoryOrder {
	return &CollectionOrder{
		collection: pd.CreateCollection(),
	}
}

// CreateOrder implements domain.RepositoryOrder.
func (c *CollectionOrder) CreateOrder(ctx context.Context, tx *gorm.DB, order *domain.Order) error {
	result := tx.Create(order)
	return result.Error
}

// DeleteOrder implements domain.RepositoryOrder.
func (c *CollectionOrder) DeleteOrder(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetOrdersByUserId implements domain.RepositoryOrder.
func (c *CollectionOrder) GetOrdersByUserId(ctx context.Context, creator_id int64) (*domain.Order, error) {
	panic("unimplemented")
}

// UpdateOrder implements domain.RepositoryOrder.
func (c *CollectionOrder) UpdateOrder(ctx context.Context, order *domain.Order) error {
	panic("unimplemented")
}
