package conn

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"myapp/config"
)

var PostgresDatabase *gorm.DB
var e error

func ConnectPostgres() {
	pgConf := config.Postgres()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", pgConf.Host, pgConf.User, pgConf.Password, pgConf.DBName, pgConf.Port)
	PostgresDatabase, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		log.Fatal("failed to connect to postgres at ", pgConf.Host, ":", pgConf.Port)
		panic(e)
	}
	log.Info("connected to postgres at ", pgConf.Host, ":", pgConf.Port)
}

func NewPostgresClient() *gorm.DB {
	return PostgresDatabase
}
