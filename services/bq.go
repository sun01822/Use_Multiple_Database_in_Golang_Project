package services

import (
	"myapp/domain"
)

type BQService struct {
	bqRepo domain.BQRepository
}

func NewBQService(bqRepo domain.BQRepository) domain.BQUseCase {
	return BQService{
		bqRepo: bqRepo,
	}
}

func (s BQService) Get() (domain.DataDetails, error) {
	return s.bqRepo.GetFromBQ()
}
