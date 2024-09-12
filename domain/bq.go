package domain

type BQRepository interface {
	GetFromBQ(limit int64) (DataDetailsBQ, error)
}

type BQUseCase interface {
	GetDataBQ(limit int64) (DataDetailsBQ, error)
}
