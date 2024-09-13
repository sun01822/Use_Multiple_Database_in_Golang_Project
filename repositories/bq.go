package repositories

import (
	"fmt"
	"log"
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

func (bqRepo *BQRepository) GetFromBQ(payload domain.Payload) (domain.DataDetailsBQ, error) {
	// Your code here
	var resp domain.DataDetailsBQ
	var data []domain.SubmissionDataModelBQ
	conf := config.BQ()
	query := "SELECT * FROM " + conf.ProjectID + "." + conf.DatasetID + "." + conf.TableID
	if payload.UUID != "" {
		query += " WHERE uuid = '" + payload.UUID + "'"
	}
	if payload.UUID == "" && payload.Date != "" {
		query += " WHERE date = '" + payload.Date + "'"
	}
	if payload.UUID == "" && payload.Date == "" {
		query += " LIMIT " + payload.Limit
	}
	log.Println(query)
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
