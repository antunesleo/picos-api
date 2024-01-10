package domain

import (
	"github.com/antunesleo/picos-api/spots/infrastructure/transactions"
)

type SpotRepository interface {
	ListAll(tx *transactions.Transaction) ([]Spot, error)
}
