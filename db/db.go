package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormConfig struct {
}

func NewGormDB(dsn string) *gorm.DB {
	gormDialector := mysql.Open(dsn)
	gormDb, err := gorm.Open(gormDialector)
	if err != nil {
		fmt.Println(err)
	}
	return gormDb
}
