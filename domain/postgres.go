package domain

type PostRepository interface {
	GetFromPostgres(limit int64) (DataDetailsPG, error)
}

type PostUseCase interface {
	GetFromPG(limit int64) (DataDetailsPG, error)
}
