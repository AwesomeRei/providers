package messages

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"time"
)

type Type int




const (
	UnknownType Type = iota
	ExampleMessageType
	GenericMessageType
)

type ExampleMessage struct {
	Name string `json:"name"`
	Time time.Time `json:"time"`
	Message string `json:"message"`
}

type GenericMessage struct {
	Message string `json:"message"`
}

func ToType(str string) Type  {
	switch str {
	case "ExampleMessage":
		return ExampleMessageType
	case "GenericMessage":
		return GenericMessageType
	default:
		return UnknownType
	}
}

func (t Type) String () string {
	switch t {
	case ExampleMessageType:
		return "ExampleMessageType"
	case GenericMessageType:
		return "GenericMessageType"
	default:
		return "UnknownType"
	}
}

func (t Type) ValidateStruct(i interface{}) bool{
	var valid bool
	switch t {
	case ExampleMessageType:
		_,valid = i.(ExampleMessage)
	case GenericMessageType:
		_,valid = i.(GenericMessage)

	default:
		valid = false
	}
	return valid
}

func (t Type) GetStruct(data []byte) (interface{},error) {
	switch t {
	case ExampleMessageType:
		var i ExampleMessage
		err := Deserialize(data,&i)
		return i,err
	case GenericMessageType:
		var i GenericMessage
		err := Deserialize(data,&i)
		return i,err
	default:
		return nil,errors.New("unsupported Message Type")
	}
}


func Serialize(i interface{}) ([]byte,error)  {
	b,err := json.Marshal(i)
	if err != nil{
		log.Error().Str("Serialize Message",err.Error()).Send()
		return nil, errors.New("error serializing message")
	}
	return b,nil
}

func Deserialize(data []byte,i interface{}) error  {
	err := json.Unmarshal(data,i)
	if err != nil{
		log.Error().Str("Deserialize Message",err.Error()).Send()
		return errors.New("error deserializing message")
	}
	return nil
}