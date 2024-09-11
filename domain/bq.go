package domain

type Repository interface {
	GetFromBQ() (DataDetails, error)
}

type UseCase interface {
	Get() (DataDetails, error)
}
