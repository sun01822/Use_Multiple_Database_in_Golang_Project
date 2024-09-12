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

func (s BQService) GetDataBQ(limit int64) (domain.DataDetailsBQ, error) {
	return s.bqRepo.GetFromBQ(limit)
}
