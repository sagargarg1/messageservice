package model

import (
	"fmt"
)

// ErrMessageNotFound is an error raised when a message can not be found in the database
var (
	ErrMessageNotFound = fmt.Errorf("Message not found")
)

// Message defines the structure for an API message
type Message struct {
        ID int `json:"id"`
        Text string `json:"message" validate:"required"`
}

// GenericError is a generic error message returned by a server
type GenericError struct {
        Message string `json:"message"`
}

var Metrics map[string]int = map[string]int{
	"Number":                   0,
	"BirthdayMessages":         0,
	"SorryMessages":            0,
	"GoodMorningMessages":      0,
	"PalindromeMessages":       0,
}
