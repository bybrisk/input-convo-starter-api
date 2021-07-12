package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/input-convo-starter-api/data"
)

func TestGetInitialConversationCRUDOPS(t *testing.T) {

	payload := &data.InitialConversationRequest{
		UserID: "60b246f9ff392f363c406fd3",
		BusinessID: "TESTID",
	}

	res:= data.GetInitialConversationCRUDOPS(payload) 

	fmt.Println(res)
}

/*func TestGetActionHandlerQuestion(t *testing.T) {

	payload := &data.ActionHandlerRequest{
		ActionHandler:"feedback",
		UserID: "60c4fbce05f858fc52b64ea5",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
	}

	res:= data.GetActionHandlerQuestion(payload) 

	fmt.Println(res)
}*/
