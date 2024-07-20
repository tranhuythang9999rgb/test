package domain

import (
	"context"

	"gorm.io/gorm"
)

// Product represents a product in the database
type Product struct {
	ID              int64   `json:"id"`
	CreatorID       int64   `json:"creator_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Quantity        int     `json:"quantity"`
	Description     string  `json:"description"`
	DiscountPercent float64 `json:"discount_percent"`
	StatusSell      bool    `json:"status_sell"`
	IsActive        bool    `json:"is_active"`
	CreateTime      int64   `json:"create_time"`
	UpdateTime      int64   `json:"update_time"`
}

// ProductService defines the interface for product operations
type RepositoryProducts interface {
	CreateProduct(ctx context.Context, product *Product) error
	GetProductByID(ctx context.Context, id int64) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id int64) error
	ListProducts(ctx context.Context) ([]*Product, error)
	UpdateProductQuantityById(ctx context.Context, tx *gorm.DB, id int64, quantity int) error
}
