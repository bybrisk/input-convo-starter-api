
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-convo-starter-api/data"
)

// swagger:route POST /commence/get commence startAConversation
// Get the conversation starting data for a business.
//
// responses:
//	200: intialConversationResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Convo) Get_Conversation_Initials (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-convo-starter-api Module")
	request := &data.InitialConversationRequest{}

	err:=request.FromJSONToInitialConversationRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = request.ValidateInitialConversationRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-convo-starter-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//get data from mongo
	response := data.GetInitialConversationCRUDOPS(request)

	//writing to the io.Writer
	w.Header().Set("Content-Type", "application/json")
	
	err = response.InitialConversationResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}