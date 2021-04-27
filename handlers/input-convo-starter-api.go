package handlers

import (
	"log"
)

type Input_Convo struct {
 l *log.Logger
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func New_Input_Convo(l *log.Logger) *Input_Convo{
	return &Input_Convo{l}
}