package services

import (
	"myapp/domain"
)

type BQService struct {
	bqRepo domain.Repository
}

func NewBQService(bqRepo domain.Repository) domain.UseCase {
	return BQService{
		bqRepo: bqRepo,
	}
}

func (s BQService) Get() (domain.DataDetails, error) {
	return s.bqRepo.GetFromBQ()
}
