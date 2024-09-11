package services

import "myapp/domain"

type PostgresService struct {
	postgresRepo domain.PostRepository
}

func NewPostgresService(postgresRepo domain.PostRepository) domain.PostUseCase {
	return PostgresService{
		postgresRepo: postgresRepo,
	}
}

func (p PostgresService) Get() (domain.DataDetails, error) {
	return p.postgresRepo.GetFromPostgres()
}
