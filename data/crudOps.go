package data

func GetInitialConversationCRUDOPS(d *InitialConversationRequest) *InitialConversationResponse{
	//isUserRegistered

	var response InitialConversationResponse

	var actionHandlerArray []string

	resp,err:= GetInitialConversationMongo(d.BusinessID)
	if err!=nil{
		response= InitialConversationResponse{
			Message:"",
			ActionHandlers:actionHandlerArray,
			Response:"Error! Some error occured!",
		}
	} else{
		response= InitialConversationResponse{
			Message:resp.Message,
			ActionHandlers:resp.ActionHandlers,
			Response:"success",
		}
	}
	return &response
}