package pgsql

import (
	"ap_sell_products/common/configs"
	"ap_sell_products/core/domain"
	"ap_sell_products/core/infra"
	"context"

	"gorm.io/gorm"
)

type CollectionUser struct {
	collection *gorm.DB
}

func NewCollectionUser(cf *configs.Configs, user *infra.PostGresql) domain.RepositoryUser {
	return &CollectionUser{
		collection: user.CreateCollection(),
	}
}

func (c *CollectionUser) AddUser(ctx context.Context, req *domain.User) error {
	result := c.collection.Create(req)
	return result.Error
}

func (c *CollectionUser) FindUserByUserName(ctx context.Context, userName string) (*domain.User, error) {
	var user *domain.User
	result := c.collection.Where("user_name = ? ", userName).First(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return user, result.Error
}
