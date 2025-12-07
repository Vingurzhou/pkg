package db

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConfig struct {
}
type InfluxdbConfig struct {
}

func NewGormDB(dialector gorm.Dialector) *gorm.DB {
	gormDb, err := gorm.Open(dialector, &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	return gormDb
}

func NewInfluxDBCli(token, url string) influxdb2.Client {
	client := influxdb2.NewClient(url, token)
	return client
}
