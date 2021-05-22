
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-convo-starter-api/data"
)

// swagger:route POST /commence/action commence getActionHandlers
// Get the action handlers for the business.
//
// responses:
//	200: actionHandlersResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Convo) Get_Action_Handlers (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-convo-starter-api Module")
	request := &data.ActionHandlerRequest{}

	err:=request.FromJSONToActionHandlerRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = request.ValidateActionHandlerRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-convo-starter-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add delivery to elastic search
	response := data.UpdateDeliveryStatusCO(request)

	//writing to the io.Writer
	err = response.ActionHandlerResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}