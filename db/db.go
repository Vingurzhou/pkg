package db

import (
	"fmt"

	"gorm.io/gorm"
)

type GormConfig struct {
}

func NewGormDB(dialector gorm.Dialector) *gorm.DB {
	gormDb, err := gorm.Open(dialector, &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	return gormDb
}
