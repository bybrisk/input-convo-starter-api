package data

func GetInitialConversationCRUDOPS(d *InitialConversationRequest) *InitialConversationResponse{
	
	var response InitialConversationResponse
	var actionHandlerArray []ActionHandlerObject
	var emptyCardObject cardObject

	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= InitialConversationResponse{
			ActionHandlers:actionHandlerArray,
			QCard: emptyCardObject,
			Response:"Error! User not registered!",
			Status:403,
		}
	} else{
		resp,err:= GetInitialConversationMongo(d.BusinessID)
		if err!=nil{
			response= InitialConversationResponse{
				ActionHandlers:actionHandlerArray,
				QCard: emptyCardObject,
				Response:"Error! Some error occured!",
				Status:501,
			}
		} else{
			response= InitialConversationResponse{
				ActionHandlers:resp.ActionHandlers,
				QCard: resp.QCard,
				Response:"successfully found intial conversation!",
				Status:200,
			}
		}
	}
	return &response
}

func GetActionHandlerQuestion (d *ActionHandlerRequest) *ActionHandlerResponse{
	var response ActionHandlerResponse
	var emptyQuestionObjectArray []QuestionObjectArray
	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= ActionHandlerResponse{
			QuestionArray: emptyQuestionObjectArray,
			ActionHandler: d.ActionHandler,
			Response:"Error! User not registered!",
			Status:403,
		}
	} else{
		resp,err:= GetQuestionForActionHandlerMongo(d)
		if err!=nil{
			response= ActionHandlerResponse{
				QuestionArray: resp ,
				ActionHandler: d.ActionHandler,
				Response:"Error! Some error occured!",
				Status:501,
			}
		} else{
			response= ActionHandlerResponse{
				QuestionArray: resp ,
				ActionHandler: d.ActionHandler,
				Response:"success",
				Status:200,
			}
		}
	}

	return &response
}