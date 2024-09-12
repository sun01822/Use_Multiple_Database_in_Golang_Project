package repositories

import (
	"fmt"
	"myapp/config"
	"myapp/conn"
	"myapp/domain"
	"time"
)

type BQRepository struct {
	bqClient *conn.MyAppBQDatabase
}

func NewBQRepository(bqClient *conn.MyAppBQDatabase) BQRepository {
	return BQRepository{
		bqClient: bqClient,
	}
}

func (bqRepo *BQRepository) GetFromBQ(limit int64) (domain.DataDetailsBQ, error) {
	// Your code here
	var resp domain.DataDetailsBQ
	var data []domain.SubmissionDataModelBQ
	conf := config.BQ()
	query := "SELECT * FROM " + conf.ProjectID + "." + conf.DatasetID + "." + conf.TableID + " LIMIT " + fmt.Sprintf("%d", limit)

	startTime := time.Now()
	if err := bqRepo.bqClient.Get(&domain.SubmissionDataModelBQ{}, query, &data); err != nil {
		return domain.DataDetailsBQ{}, err
	}
	durationTime := time.Since(startTime)

	hours := int(durationTime.Hours())
	minutes := int(durationTime.Minutes()) % 60
	seconds := int(durationTime.Seconds()) % 60
	resp.FetchingTime = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	resp.Data = data
	return resp, nil
}
