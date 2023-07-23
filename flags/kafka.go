package flags

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type KafkaCFG struct {
	Brokers    string `json:"brokers"`
	Topic      string `json:"topic"`
	Group      string `json:"group"`
	Partitions int    `json:"partitions"`
}

func (k KafkaCFG) Validate() error {
	if k.Brokers == "" {
		return errors.New("'brokers' has been missing in KafkaCFG")
	}
	if k.Topic == "" {
		return errors.New("'topic' has been missing in KafkaCFG")
	}
	if k.Partitions <= 0 {
		return errors.New("'partitions' has been missing in KafkaCFG")
	}
	return nil
}

func DefineKafkaFlags(cfg *KafkaCFG, cmd, tag string) {
	flag.StringVar(&cfg.Brokers,
		fmt.Sprintf("%sbrokers", tag),
		"localhost:9092",
		fmt.Sprintf(strings.ReplaceAll("Set the KafkaComm brokers\n\t %s --{tag}comm kafka --{tag}brokers 'kafka-0:9092,kafka-1:9092,kafka-02:9092' --{tag}topic 'topic-1' --{tag}partitions 3", "{tag}", tag), cmd))
	flag.StringVar(&cfg.Topic,
		fmt.Sprintf("%stopic", tag),
		"",
		fmt.Sprintf(strings.ReplaceAll("Set the KafkaComm topic\n\t %s --{tag}comm kafka --{tag}brokers 'kafka-0:9092,kafka-1:9092,kafka-02:9092' --{tag}topic 'topic-1' --{tag}partitions 3", "{tag}", tag), cmd))
	flag.StringVar(&cfg.Group,
		fmt.Sprintf("%sgroup", tag),
		fmt.Sprintf("%s%d", tag, time.Now().UnixNano()),
		fmt.Sprintf(strings.ReplaceAll("Set the KafkaComm group\n\t %s --{tag}comm kafka --{tag}brokers 'kafka-0:9092,kafka-1:9092,kafka-02:9092' --{tag}topic 'topic-1' --{tag}partitions 3 --{tag}group 'g1'", "{tag}", tag), cmd))
	flag.IntVar(&cfg.Partitions,
		fmt.Sprintf("%spartitions", tag),
		3,
		fmt.Sprintf(strings.ReplaceAll("Set the KafkaComm partitions\n\t %s --{tag}comm kafka --{tag}brokers 'kafka-0:9092,kafka-1:9092,kafka-02:9092' --{tag}topic 'topic-1' --{tag}partitions 3", "{tag}", tag), cmd))
}
