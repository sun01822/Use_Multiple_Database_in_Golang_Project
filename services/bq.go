package services

import "myapp/repositories"

type BQService struct {
	bqRepo repositories.BQRepository
}

func NewBQService(bqRepo repositories.BQRepository) BQService {
	return BQService{
		bqRepo: bqRepo,
	}
}
