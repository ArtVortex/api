package pg

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type PostgresAdapter struct {
	DB *gorm.DB
}

func NewPGAdapter(dsn string) (*PostgresAdapter, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to Postgres", err)
		return nil, err
	}

	Migrate(db)

	DB = db

	sqlDB, _ := db.DB()

	// Set the maximum connection lifetime to 5 minutes
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return &PostgresAdapter{DB: db}, nil
}

func (p *PostgresAdapter) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *PostgresAdapter) Ping() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (p *PostgresAdapter) GetDB() *gorm.DB {
	return p.DB
}
