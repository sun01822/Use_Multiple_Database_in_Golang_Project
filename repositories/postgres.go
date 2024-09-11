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

func (p PostgresRepository) GetFromPostgres() (domain.DataDetails, error) {
	pgConf := config.Postgres()

	var resp domain.DataDetails
	var data []domain.SubmissionDataModel

	startTime := time.Now()
	if err := p.postgresClient.Table(pgConf.TableName).Find(&data).Error; err != nil {
		return domain.DataDetails{}, err
	}
	durationTime := time.Since(startTime)

	hours := int(durationTime.Hours())
	minutes := int(durationTime.Minutes()) % 60
	seconds := int(durationTime.Seconds()) % 60
	resp.FetchingTime = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	resp.Data = data
	return resp, nil

}
