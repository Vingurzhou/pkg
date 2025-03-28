package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewGormDB(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	gormDB := NewGormDB(os.Getenv("DSN"), GormConfig{})
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
