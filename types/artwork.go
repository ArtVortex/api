package types

import (
	"time"

	"github.com/google/uuid"
)

type ArtWork struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	NftId       string
	Name        string
	Description string
	Creator     string
	ImageUrl    string
	CreditCost  int
	Prompt      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
