package db

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConfig struct {
}

func NewGormDB(dialector gorm.Dialector) *gorm.DB {
	gormDb, err := gorm.Open(dialector, &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
	}
	return gormDb
}
