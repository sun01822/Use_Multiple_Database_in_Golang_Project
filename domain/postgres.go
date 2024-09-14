package domain

type PostRepository interface {
	GetFromPostgres(payload Payload) (DataDetailsPG, error)
}

type PostUseCase interface {
	GetFromPG(payload Payload) (DataDetailsPG, error)
}
