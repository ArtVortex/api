package types

import (
	"time"
)

type NFT struct {
	UID           uint `gorm:"primary_key"`
	ArtWorkId     string
	CollectionId  int
	Name          string
	Description   string
	Metadata      map[string]interface{} `gorm:"type:jsonb"`
	Creator       string                 // only have value when NFT is mint from our platform
	Owner         string                 //wallet_address
	TokenAddress  string
	TokenType     string `gorm:"default:erc721"` //erc721 or erc1155
	ContentUrl    string
	BlockChainId  int `gorm:"default:137"` //eth = 1, polygon = 137
	Price         int
	TradeType     string `gorm:"default:fixed"` //acution/fixed
	ListingStatus int    //0/1
	Currency      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
