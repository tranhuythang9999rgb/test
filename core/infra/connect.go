package infra

import (
	"ap_sell_products/common/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostGresql struct {
	config   *configs.Configs
	postgres *gorm.DB
}

func NewpostgreDb(config *configs.Configs) *PostGresql {

	postgresDb := &PostGresql{
		config: config,
	}
	err := postgresDb.connect()
	log.Println("starting connect postgre")
	if err != nil {
		log.Fatalf("fail to connect postgre", err)

	}
	return postgresDb
}

func (p *PostGresql) connect() error {
	dsn := configs.Get().DataSource
	log.Println("error ", dsn)
	if dsn == "" {
		return fmt.Errorf("empty postgre uri")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgre: %w", err)
	}

	p.postgres = db
	log.Println("Connected to postgre successfully")
	return nil
}
func (p *PostGresql) CreateCollection() *gorm.DB {
	return p.postgres
}
