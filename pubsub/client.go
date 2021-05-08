package pubsub

import (
	"errors"

)

type MessageBox struct {
	ID          string
	Data        []byte
	MessageType messages.Type
}

func NewMessageBox(id string ,message interface{},t messages.Type) (MessageBox,error){
	data,err := messages.Serialize(message)
	if err != nil {
		return MessageBox{},err
	}
	if !t.ValidateStruct(message) {
		return MessageBox{}, errors.New("invalid struct")
	}
	return MessageBox{
		Data: data,
		MessageType: t,
		ID: id,
	}, err
}