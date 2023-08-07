package types

import (
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string
	Description  string
	Creator      string
	Url          string
	TokenAddress string
	BlockChain   string
	NumberOfNfts int16
	CreatedAt    time.Time
}
