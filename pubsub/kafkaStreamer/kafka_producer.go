package kafkaStreamer

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/AwesomeRei/kraft-producer/provider"
	"gitlab.com/AwesomeRei/kraft-producer/provider/messages"
)

type KafkaProducer struct {
	W kafka.Writer
}

type Options struct {
	Balancer kafka.Balancer
	MaxAttemps int
	RequiredAcks kafka.RequiredAcks
	Async bool
	Completion func(messages []kafka.Message, err error)
	Compression kafka.Compression
	Transport kafka.RoundTripper
}

const (
	MISSING_HOST = "host could not be nil"
)

func New(host string,port int, o Options) (*KafkaProducer,error){
	switch  {
	case host == "":
		return nil,errors.New(MISSING_HOST)
	}
	address := fmt.Sprintf("%s:%d",host,port)
	return &KafkaProducer{
		W: kafka.Writer{Addr: kafka.TCP(address),
			Balancer: o.Balancer,
			MaxAttempts: o.MaxAttemps,
			RequiredAcks: o.RequiredAcks,
			Async: o.Async,
			Completion: o.Completion,
			Compression: o.Compression,
			Transport: o.Transport,
		},
	},nil
}

func (k *KafkaProducer) SendEvent(ctx context.Context,topic string,msg provider.MessageBox) error {
	b,err := messages.Serialize(msg)
	if err != nil{
		return err
	}
	err = k.W.WriteMessages(ctx,kafka.Message{
		Topic: topic,
		Key: []byte(topic),
		Value: b,
	})
	if err != nil {
		return err
	}
	return nil
}