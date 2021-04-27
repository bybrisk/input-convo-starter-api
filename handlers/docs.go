// Package classification of Input-Convo-Starter API
//
// Documentation for Input-Convo-Starter API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/input-convo-starter-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Message on fetching intial conversation for a business
// swagger:response intialConversationResponse
type intialConvoPostResponseWrapper struct {
	// response on getting intial conversation for a business
	// in: body
	Body data.InitialConversationResponse
}

// swagger:parameters startAConversation
type startConvoParmsWrapper struct {
	// Data structure to get intial conversation of any business
	// in: body
	// required: true
	Body data.InitialConversationRequest
}