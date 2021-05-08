package messages

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestSerialize(t *testing.T) {
	tests := []struct{
		name string
		i interface{}
		t Type
		err error
	}{
		{
			name: "Serialize Success",
			i: ExampleMessage{
				Name: "name",
				Time: time.Now(),
				Message: "message",
			},
			t: ExampleMessageType,
			err: nil,
		},
		{
			name: "Serialize Error",
			i: math.Inf(1),
			t: ExampleMessageType,
			err: errors.New("error serializing message"),
		},

	}
	for _,testCase := range tests{
		t.Run(testCase.name, func(t *testing.T) {
			 _,err := Serialize(testCase.i)
			 assert.Equal(t, testCase.err,err)
		})
	}
}

func TestDeserialize(t *testing.T) {
	tests := []struct{
		name string
		i interface{}
		expectedResult interface{}
		t Type
		data []byte
		err error
	}{
		{
			name: "Deserialize Success",
			i: ExampleMessageType,
			expectedResult: ExampleMessage{
				Name: "name",
				Message: "message",
			},
			data: []byte(" {\"Name\":\"name\",\"Message\":\"message\"}"),
			t: ExampleMessageType,
			err: nil,
		},
		{
			name: "Deserialize General Message Success",
			i: GenericMessage{},
			expectedResult: ExampleMessage{
				Name: "name",
				Message: "message",
			},
			data: []byte("{ \"message\": \"message\"}"),
			t: GenericMessageType,
			err: nil,
		},
	}
	for _,testCase := range tests{
		t.Run(testCase.name, func(t *testing.T) {
			err := Deserialize(testCase.data,&testCase.i)
			t.Log(testCase.i)
			//assert.Equal(t, testCase.expectedResult,testCase.i)
			assert.Equal(t, testCase.err,err)
		})
	}
}