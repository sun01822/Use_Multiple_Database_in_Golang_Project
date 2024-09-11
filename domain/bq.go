package domain

type BQRepository interface {
	GetFromBQ() (DataDetails, error)
}

type BQUseCase interface {
	Get() (DataDetails, error)
}
