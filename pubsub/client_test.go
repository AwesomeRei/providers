package pubsub

import (
	"errors"
	"github.com/AwesomeRei/providers/pubsub/messages"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestNewMessageBox(t *testing.T) {
	tests := []struct{
		name string
		id string
		message interface{}
		t messages.Type
		expected error
	}{
		{
			name: "success create new message box",
			id: "id",
			message: messages.ExampleMessage{
				Name:    "name",
				Time:    time.Now(),
				Message: "message",
			},
			t: messages.ExampleMessageType,
			expected: nil,
		},
		{
			name: "error validate struct",
			id: "id",
			message: "string",
			t: messages.ExampleMessageType,
			expected: errors.New("invalid struct"),
		},
		{
			name: "error serialize message",
			id: "id",
			message: math.Inf(1),
			t: messages.ExampleMessageType,
			expected: errors.New("error serializing message"),
		},
	}

	for _,testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_,err := NewMessageBox(testCase.id,testCase.message,testCase.t)
			assert.Equal(t, testCase.expected,err)
		})
	}
}
