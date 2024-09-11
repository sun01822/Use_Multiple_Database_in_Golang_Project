package conn

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/labstack/gommon/log"
	"myapp/config"
)

var bqDatabase MyAppBQDatabase

type MyAppBQDatabase struct {
	BqDB *bigquery.Client
}

func ConnectBQ() {
	conf := config.BQ()
	ctx := context.Background()

	log.Info("connecting to bigquery at ", conf.ProjectID, ":", conf.DatasetID, "...")

	client, clientErr := bigquery.NewClient(ctx, conf.ProjectID)
	if clientErr != nil {
		log.Fatalf("bigquery.NewClient: %v", clientErr)
	}
	log.Info("connected to bigquery at ", conf.ProjectID, ":", conf.DatasetID, "...")

	defer client.Close()

	bqDatabase.BqDB = client
}

func NewBqDBClient() *MyAppBQDatabase {
	return &bqDatabase
}
