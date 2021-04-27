package data

import (
	"github.com/go-playground/validator/v10"
)

//post request for registering a user
type InitialConversationResponse struct{
	// Message the business wants the user to see.
	//
	Message string `json: "message"`

	// The action handler which decides which API to fire
	//
	ActionHandlers []string `json: "actionHandlers"`

	//Response message
	Response string `json:"response"`
}

//post request for getting the inital conversation of a business
type InitialConversationRequest struct{
	// UserID of the user 
	//
	// required: true
	// max length: 1000
	UserID string `json: "userID" validate:"required"`

	// BusinessID of the business user is subscribing to
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessid" validate:"required"`
}

func (d *InitialConversationRequest) ValidateInitialConversationRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}