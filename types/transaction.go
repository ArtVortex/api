package types

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	HashId     string    //transaction_id from block chain
	NftId      string
	ArtWorkId  string
	Buyer      string
	Type       string //create, transfer or mint
	CurrencyID string
	Status     string //reject, approve, completed, failed...
	CreatedAt  time.Time
}
