package domain

import (
	"gorm.io/gorm"
)

type RepositoryTrans interface {
	WithTransaction(fn func(*gorm.DB) error) error
}
