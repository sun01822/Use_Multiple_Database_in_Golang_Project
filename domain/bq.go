package domain

type BQRepository interface {
	GetFromBQ(payload Payload) (DataDetailsBQ, error)
}

type BQUseCase interface {
	GetDataBQ(payload Payload) (DataDetailsBQ, error)
}
