package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"myapp/config"
	"myapp/domain"
	"time"
)

type PostgresRepository struct {
	postgresClient *gorm.DB
}

func NewPostgresRepository(postgresClient *gorm.DB) *PostgresRepository {
	return &PostgresRepository{postgresClient: postgresClient}
}

func (p PostgresRepository) GetFromPostgres(limit int64) (domain.DataDetailsPG, error) {
	pgConf := config.Postgres()
	convertLimit := int(limit)

	var resp domain.DataDetailsPG
	var data []domain.SubmissionDataModelPG

	startTime := time.Now()
	if err := p.postgresClient.Table(pgConf.TableName).Limit(convertLimit).Find(&data).Error; err != nil {
		return domain.DataDetailsPG{}, err
	}
	durationTime := time.Since(startTime)

	hours := int(durationTime.Hours())
	minutes := int(durationTime.Minutes()) % 60
	seconds := int(durationTime.Seconds()) % 60
	resp.FetchingTime = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	resp.Data = data
	return resp, nil
}
