package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"log"
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

func (p PostgresRepository) GetFromPostgres(payload domain.Payload) (domain.DataDetailsPG, error) {
	pgConf := config.Postgres()

	query := "SELECT * FROM " + pgConf.TableName
	if payload.UUID != "" {
		query += " WHERE uuid = '" + payload.UUID + "'"
	}
	if payload.UUID == "" && payload.Date != "" {
		query += " WHERE date = '" + payload.Date + "'"
	}
	if payload.UUID == "" && payload.Date == "" {
		query += " LIMIT " + payload.Limit
	}

	var resp domain.DataDetailsPG
	var data []domain.SubmissionDataModelPG
	log.Println(query)
	startTime := time.Now()
	if err := p.postgresClient.Raw(query).Scan(&data).Error; err != nil {
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
