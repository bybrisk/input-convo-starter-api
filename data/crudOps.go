package data

func GetInitialConversationCRUDOPS(d *InitialConversationRequest) *InitialConversationResponse{
	
	var response InitialConversationResponse
	var actionHandlerArray []string

	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= InitialConversationResponse{
			Message:"",
			ActionHandlers:actionHandlerArray,
			Response:"Error! User not registered!",
		}
	} else{
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
	}
	return &response
}

func GetActionHandlerQuestion (d *ActionHandlerRequest) *ActionHandlerResponse{
	var response ActionHandlerResponse
	var EmptyArr []string
	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= ActionHandlerResponse{
			Questions: EmptyArr,
			QType: EmptyArr,
			ActionHandler: d.ActionHandler,
			Response:"Error! User not registered!",
		}
	} else{
		resp,err:= GetQuestionForActionHandlerMongo(d)
		if err!=nil{
			response= ActionHandlerResponse{
				Questions: resp.Questions ,
				QType: resp.QType ,
				ActionHandler: d.ActionHandler,
				Response:"Error! Some error occured!",
			}
		} else{
			response= ActionHandlerResponse{
				Questions: resp.Questions ,
				QType: resp.QType ,
				ActionHandler: d.ActionHandler,
				Response:"success",
			}
		}
	}

	return &response
}