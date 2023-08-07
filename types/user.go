package types

import (
	"time"
)

type User struct {
	ID            uint
	UserName      string `gorm:"uniqueIndex"`
	Email         string `gorm:"uniqueIndex"`
	CreditBalance int    `gorm:"default:0"`
	DisplayName   string
	WalletAddress string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
