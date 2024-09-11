package repositories

import "myapp/conn"

type BQRepository struct {
	bqClient *conn.MyAppBQDatabase
}

func NewBQRepository(bqClient *conn.MyAppBQDatabase) BQRepository {
	return BQRepository{
		bqClient: bqClient,
	}
}
