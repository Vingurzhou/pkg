package mq

import (
	"crypto/tls"
	"errors"
	"strings"

	"github.com/IBM/sarama"
)

type SaramaConfig struct {
	Brokers    string
	Username   string
	Password   string
	TLSEnabled bool
	Algorithm  string
	Assignor   string
}

// TODO
func createTLSConfiguration() (t *tls.Config) {
	panic("not implement")
}
func NewSaramCli(k SaramaConfig) (sarama.Client, error) {
	if k.Brokers == "" {
		return nil, errors.New("brokers is required")
	}
	splitBrokers := strings.Split(k.Brokers, ",")

	// TODO
	version, err := sarama.ParseKafkaVersion(sarama.DefaultVersion.String())
	if err != nil {
		return nil, err
	}

	if k.Username == "" {
		return nil, errors.New("username is required")
	}

	if k.Password == "" {
		return nil, errors.New("password is required")
	}

	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Version = version
	conf.ClientID = "sasl_scram_client"
	conf.Metadata.Full = true
	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = k.Username
	conf.Net.SASL.Password = k.Password
	conf.Net.SASL.Handshake = true
	if k.Algorithm == "sha512" {
		conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
		conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
	} else if k.Algorithm == "sha256" {
		conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
		conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

	} else {
		return nil, errors.New("algorithm is required")
	}

	if k.TLSEnabled {
		conf.Net.TLS.Enable = true
		conf.Net.TLS.Config = createTLSConfiguration()
	}

	switch k.Assignor {
	case "sticky":
		conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		return nil, errors.New("Unrecognized consumer group partition assignor" + k.Assignor)
	}
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Consumer.Return.Errors = true
	cli, err := sarama.NewClient(splitBrokers, conf)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
