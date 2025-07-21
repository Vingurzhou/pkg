package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	// "gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
)

func TestNewGormDB(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	// gormDB := NewGormDB(sqlite.Open(os.Getenv("DSN")))
	gormDB := NewGormDB(mysql.Open(os.Getenv("DSN")))
	db, err := gormDB.DB()
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	stats := db.Stats()
	t.Logf("%+v", stats)
}
