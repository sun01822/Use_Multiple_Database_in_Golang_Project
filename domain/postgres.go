package domain

type PostRepository interface {
	GetFromPostgres() (DataDetails, error)
}

type PostUseCase interface {
	Get() (DataDetails, error)
}
