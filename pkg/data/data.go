package data

import (
        "github.com/sagargarg1/messageservice/pkg/model"
)

var (
	MessageInterface MessageDBInterface = &MessagesDB{}
)

type MessagesDB struct {}

type MessageDBInterface interface {
	UpdateMessage(message model.Message) error
	GetAllMessages() []*model.Message
	GetMessage(id int) (*model.Message, error)
	AddMessage(message model.Message)
	DeleteMessage(id int) error
}

func (m *MessagesDB) UpdateMessage(message model.Message) error {
	i := findIndexByMessageID(message.ID)
	if i == -1 {
		return model.ErrMessageNotFound
	}

	messageList[i] = &message

	return nil
}

func (m *MessagesDB) GetAllMessages() []*model.Message {

	return messageList
}

func (m *MessagesDB) GetMessage(id int) (*model.Message, error) {
	i := findIndexByMessageID(id)
	if i == -1 {
		return nil, model.ErrMessageNotFound
	}

	message := *messageList[i]
	return &message, nil
}

func (m *MessagesDB) AddMessage(message model.Message) {
	maxID := messageList[len(messageList)-1].ID
	message.ID = maxID + 1
	messageList = append(messageList, &message)
}

func (p *MessagesDB) DeleteMessage(id int) error {
	i := findIndexByMessageID(id)
	if i == -1 {
		return model.ErrMessageNotFound
	}

	messageList = append(messageList[:i], messageList[i+1])

	return nil
}

// findIndex finds the index of a message in the database
// returns -1 when no product can be found
func findIndexByMessageID(id int) int {
	for i, p := range messageList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var messageList = []*model.Message{
	&model.Message{
		ID:   1,
		Text: "I love India",
	},
	&model.Message{
		ID:   2,
		Text: "I love Bengaluru",
	},
}
