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
		Brokers:    "221.214.51.164:19092",
		Username:   "microsate",
		Password:   "microsate_pass",
		TLSEnabled: tlsEnabled,
		Algorithm:  "sha256",
		Assignor:   "range",
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

func TestCommunity(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	t.Log(client.Brokers())
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		t.Fatal(err)
	}

	producer.Input() <- &sarama.ProducerMessage{
		Topic: "test-topic",
		Key:   sarama.StringEncoder("key1"), // 可选
		Value: sarama.StringEncoder("hello kafka"),
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(consumer.Partitions("test-topic"))
	t.Log(consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest))
}
