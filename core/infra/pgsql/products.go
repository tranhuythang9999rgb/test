package pgsql

import (
	"ap_sell_products/core/domain"
	"ap_sell_products/core/infra"
	"context"

	"gorm.io/gorm"
)

type CollectionProduct struct {
	collection *gorm.DB
}

// NewCollectionProduct creates a new instance of CollectionProduct
func NewCollectionProduct(pd *infra.PostGresql) domain.RepositoryProducts {
	return &CollectionProduct{
		collection: pd.CreateCollection(),
	}
}

// CreateProduct implements domain.RepositoryProducts.
func (c *CollectionProduct) CreateProduct(ctx context.Context, product *domain.Product) error {
	return c.collection.WithContext(ctx).Create(product).Error
}

// DeleteProduct implements domain.RepositoryProducts.
func (c *CollectionProduct) DeleteProduct(ctx context.Context, id int64) error {
	return c.collection.WithContext(ctx).Delete(&domain.Product{}, id).Error
}

// GetProductByID implements domain.RepositoryProducts.
func (c *CollectionProduct) GetProductByID(ctx context.Context, id int64) (*domain.Product, error) {
	var product *domain.Product
	if err := c.collection.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// ListProducts implements domain.RepositoryProducts.
func (c *CollectionProduct) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := c.collection.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct implements domain.RepositoryProducts.
func (c *CollectionProduct) UpdateProduct(ctx context.Context, product *domain.Product) error {
	return c.collection.WithContext(ctx).Save(product).Error
}

func (c *CollectionProduct) UpdateProductQuantityById(ctx context.Context, tx *gorm.DB, id int64, quantity int) error {
	result := tx.Model(&domain.Product{}).Where("id = ?", id).UpdateColumn("quantity", quantity)
	return result.Error
}
