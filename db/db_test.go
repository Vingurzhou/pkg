package db

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	code := m.Run()
	os.Exit(code)
}
func TestNewGormDB(t *testing.T) {

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

func TestNewInfluxDBCli(t *testing.T) {
	client := NewInfluxDBCli(os.Getenv("INFLUXDB_TOKEN"), os.Getenv("INFLUXDB_URL"))
	//////////////////////////
	writeAPI := client.WriteAPIBlocking(os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			t.Fatal(err)
		}
	}
	//////////////////////////////////
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))
	query := `from(bucket: "mybucket")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		t.Fatal(err)
	}
	for results.Next() {
		t.Log(results.Record())
	}
	if err := results.Err(); err != nil {
		t.Fatal(err)
	}
}
