package application

import (
	"github.com/antunesleo/picos-api/spots/domain"
	"github.com/antunesleo/picos-api/spots/infrastructure/transactions"
)

type SpotUseCases interface {
	List() []domain.Spot
}

type SpotsUseCasesImpl struct {
	TransactionManager transactions.TransactionManager
	SpotRepository     domain.SpotRepository
}

func (uc *SpotsUseCasesImpl) List() ([]domain.Spot, error) {
	tx, err := uc.TransactionManager.Begin()
	if err != nil {
		return []domain.Spot{}, err
	}
	spots, err := uc.SpotRepository.ListAll(tx)
	if err != nil {
		uc.TransactionManager.Commit(tx)
		return []domain.Spot{}, err
	}

	uc.TransactionManager.Rollback(tx)
	return spots, nil
}
