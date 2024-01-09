package persistence

import (
	"github.com/antunesleo/picos-api/spots/domain"
	"github.com/jmoiron/sqlx"
)

type SQLXSpotRepository struct {
	DB *sqlx.DB
}

func (r *SQLXSpotRepository) ListAll() ([]domain.Spot, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}

	spots := []domain.Spot{}
	query := "SELECT name FROM spots;"
	rows, err := tx.Query(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		spots = append(spots, domain.Spot{Name: name})
	}

	tx.Commit()

	return spots, nil
}
