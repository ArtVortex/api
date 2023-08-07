package pg

import (
	"artvortex-api/types"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(
		&types.User{},
		&types.NFT{},
		&types.ArtWork{},
		&types.Currency{},
		&types.Collection{},
		&types.Transaction{},
	)

	return nil
}
