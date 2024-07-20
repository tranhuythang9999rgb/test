package pgsql

import (
	"ap_sell_products/core/domain"
	"ap_sell_products/core/infra"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CollectionTransaction struct {
	db *gorm.DB
}

func NewCollectionTransaction(pd *infra.PostGresql) domain.RepositoryTrans {
	return &CollectionTransaction{
		db: pd.CreateCollection(),
	}
}

func (c *CollectionTransaction) WithTransaction(fn func(*gorm.DB) error) (err error) {
	timeout := 1 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tx := c.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("could not begin transaction: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic in transaction: %v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
			if err != nil {
				err = fmt.Errorf("could not commit transaction: %w", err)
			}
		}
	}()

	err = fn(tx)
	return err
}
