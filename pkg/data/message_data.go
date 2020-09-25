package data

import (
        "github.com/sagargarg1/messageservice/pkg/model"
)

type MessageDBInterface interface {
        UpdateMessage(message model.Message) error
        GetAllMessages() []*model.Message
        GetMessage(id int) (*model.Message, error)
        AddMessage(message model.Message) int
        DeleteMessage(id int) error
}

type MessageDB struct {}

func NewMessageDB () MessageDBInterface {
	return &MessageDB{}
}

//Function to update DB
func (m *MessageDB) UpdateMessage(message model.Message) error {
	i := findIndexByMessageID(message.ID)
	if i == -1 {
		return model.ErrMessageNotFound
	}

	messageList[i] = &message

	return nil
}

//Function to get all messages from DB
func (m *MessageDB) GetAllMessages() []*model.Message {

	return messageList
}

//Function to get a message with a particular ID
func (m *MessageDB) GetMessage(id int) (*model.Message, error) {
	i := findIndexByMessageID(id)
	if i == -1 {
		return nil, model.ErrMessageNotFound
	}

	message := *messageList[i]
	return &message, nil
}

//Function to add message and return the index
func (m *MessageDB) AddMessage(message model.Message) int {
	maxID := messageList[len(messageList)-1].ID
	message.ID = maxID + 1
	messageList = append(messageList, &message)
	return message.ID
}

//Function to delete message with a given ID
func (p *MessageDB) DeleteMessage(id int) error {
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
