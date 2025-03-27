package mq

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewSaramCli(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	tlsEnabled, _ := strconv.ParseBool(os.Getenv("TLSEnabled"))
	saramaCli, err := NewSaramCli(SaramaConfig{
		Brokers:    os.Getenv("Brokers"),
		Username:   os.Getenv("Username"),
		Password:   os.Getenv("Password"),
		TLSEnabled: tlsEnabled,
		Algorithm:  os.Getenv("Algorithm"),
		Assignor:   os.Getenv("Assignor"),
	})
	if err != nil {
		t.Fatal(err)
	}
	topics, err := saramaCli.Topics()
	if err != nil {
		t.Fatal(err)
	}
	for _, topic := range topics {
		t.Log(topic)
	}
}
