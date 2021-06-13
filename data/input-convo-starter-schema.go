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

	//status code
	//
	Status int64 `json:"status"`
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

//struct for getting the action handler request of a business
type ActionHandlerRequest struct{
	// Action handler string 
	//
	// required: true
	// max length: 1000
	ActionHandler string `json: "actionHandler" validate:"required"`

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

//struct for getting the action handler response of a business
type ActionHandlerResponse struct{
	// Questions to prepare the payload 
	//
	//Questions []string `json: "questions"`

	// Question type for adjusting the UI
	//
	//QType []string `json: "qType"`

	//Question objects array
	//
	QuestionArray []QuestionObjectArray `json:"questionArray"`

	// Action handler
	//
	ActionHandler string `json: "actionHandler"`

	// API response message
	//
	Response string `json:"response"`

	//status code
	//
	Status int64 `json:"status"`
}

type QuestionAndTypeStruct struct{
	Questions []string `json:"questions"`
	QType []string `json:"qtype"` 
	QContext []string `json:qcontext`
	QCustomResponseChoice [][]string `json:"qcustomresponsechoice"`
}

type QuestionObjectArray struct {

	//Question to be displayed
	//
	Question string `json:"question"`

	//Type of response expected (needed for setting the UI)
	//
	ResponseType string `json:"responseType"`
	
	//Context of the question (helps to prepare the payload)
	//
	QuestionContext string `json:"questionContext"`

	//Choice of response provided by the business
	//
	CustomResponseChoice []string `json:"customResponseChoice"` 

	//Type of question bot asks
	//
	QuestionType string `json:"questionType"`

	//Card object for the question
	//
	QestionCard cardObject `json:"questionCard"`

	//Card object for the answer
	//
	AnswerCard []cardObject `json:"answerCard"`
}

type cardObject struct{
    //URL of the pic on the card
	//
	CardPic string `json:"cardPic"`

	//Card title
	//
	Title string `json:"title"`

	//Subtitle of the card
	//
	SubTitle string `json:"subtitle"`

	//Story on the card
	//
	Story string `json:"story"`

	//Top right data on the card
	//
	TopRightData string `json:"topRightData"`

	//Bottom Right Data of the card
	//
	BottomRightData string `json:"bottomRightData"`
}

func (d *InitialConversationRequest) ValidateInitialConversationRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *ActionHandlerRequest) ValidateActionHandlerRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}