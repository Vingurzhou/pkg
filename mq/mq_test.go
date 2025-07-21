package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/IBM/sarama"
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
		// t.Log(topic)
		_ = topic
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient("zwz", saramaCli)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = consumerGroup.Close() }()
	// Track errors
	go func() {
		for err := range consumerGroup.Errors() {
			t.Log("ERROR", err)
		}
	}()
	ctx := context.Background()
	for {
		topics := []string{"SGS.requirement.statue"}
		handler := exampleConsumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := consumerGroup.Consume(ctx, topics, handler)
		if err != nil {
			t.Fatal(err)
		}
	}
}

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var msgValue map[string]interface{}
	for msg := range claim.Messages() {
		json.Unmarshal(msg.Value, &msgValue)
		fmt.Println(msgValue["requestId"])
	}
	return nil
}
