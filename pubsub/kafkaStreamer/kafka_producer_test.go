package kafkaStreamer

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"gitlab.com/AwesomeRei/kraft-producer/provider"
	"gitlab.com/AwesomeRei/kraft-producer/provider/messages"
	"testing"
	"time"
)

func TestNewKafkaWriter(t *testing.T) {
	type result struct {
		err error
	}

	tests := []struct{
		name string
		host string
		port int
		opt Options
		expected result
	}{
		{
			name: "success creation",
			host: "localhost",
			port: 9092,
			expected: result{
				err: nil,
				},
		},
		{
			name: "failed creation, param host not supplied",
			host: "",
			port: 9092,
			expected: result{
				errors.New("host could not be nil"),

			},
		},
	}

	for _,testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_,err := New(testCase.host,testCase.port,testCase.opt)
			assert.Equal(t, testCase.expected.err,err)
		})
	}
}

func TestSendEvent(t *testing.T) {
	msg,_ := provider.NewMessageBox("id",messages.ExampleMessage{
		Name: "jaeger", Time: time.Now(), Message: "test",
	},messages.ExampleMessageType)
	tests := []struct{
		name string
		client KafkaProducer
		msg provider.MessageBox
		ctx context.Context
		topic string
		error error
	}{
		{
			"success send message",
			KafkaProducer{
				W: kafka.Writer{Addr: kafka.TCP("localhost:9092"),Balancer: &kafka.LeastBytes{}},
			},
			msg,
			context.Background(),
			"jaeger",
			nil,
		},
	}
	for _,testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.client.SendEvent(context.Background(),testCase.topic,testCase.msg)
			assert.Equal(t, testCase.error,err)
		})
	}
}