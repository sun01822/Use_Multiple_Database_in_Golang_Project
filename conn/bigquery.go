package conn

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/labstack/gommon/log"
	"google.golang.org/api/iterator"
	"myapp/config"
	"reflect"
)

var bqDatabase MyAppBQDatabase

type MyAppBQDatabase struct {
	BqDB *bigquery.Client
}

func ConnectBQ() {
	conf := config.BQ()
	ctx := context.Background()

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

func (c *MyAppBQDatabase) Get(modelType interface{}, query string, modelsPtr interface{}) error {
	ctx := context.Background()
	it, err := c.BqDB.Query(query).Read(ctx)
	if err != nil {
		return err
	}

	models := reflect.ValueOf(modelsPtr).Elem()
	for {
		modelsPtr := reflect.New(reflect.TypeOf(modelType).Elem())
		if err := it.Next(modelsPtr.Interface()); err == iterator.Done {
			break
		} else if err != nil {
			return err
		}
		models.Set(reflect.Append(models, modelsPtr.Elem()))
	}
	return nil
}
