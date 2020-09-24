package model

import (
	"fmt"
)

// ErrMessageNotFound is an error raised when a message can not be found in the database
var (
	ErrMessageNotFound = fmt.Errorf("Product not found")
)

// Message defines the structure for an API message
type Message struct {
        ID int `json:"id"`
        Text string `json:"name" validate:"required"`
}

// GenericError is a generic error message returned by a server
type GenericError struct {
        Message string `json:"message"`
}
