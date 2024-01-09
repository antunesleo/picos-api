package persistence

import (
	"github.com/antunesleo/picos-api/spots/domain"
)

// Option 1: Transaction Service Layer
// Option 2: Transaction Handler
// Encapsula acesso a dados
type SQLXSpotRepository struct{}

func (r *SQLXSpotRepository) ListAll(tx *Transaction) ([]domain.Spot, error) {
	SQLXTx, err := tx.GetSQLXTransaction()
	if err != nil {
		return nil, err
	}

	spots := []domain.Spot{}
	query := "SELECT name FROM spots;"
	rows, err := SQLXTx.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}

		spots = append(spots, domain.Spot{Name: name})
	}
	return spots, nil
}
