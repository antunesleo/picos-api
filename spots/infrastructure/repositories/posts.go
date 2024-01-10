package posts

import (
	"github.com/antunesleo/picos-api/spots/domain"
	"github.com/antunesleo/picos-api/spots/infrastructure/transactions"
)

type SQLXSpotRepository struct{}

func (r *SQLXSpotRepository) ListAll(tx *transactions.Transaction) ([]domain.Spot, error) {
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
