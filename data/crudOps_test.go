package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/input-convo-starter-api/data"
)

func TestGetInitialConversationCRUDOPS(t *testing.T) {

	payload := &data.InitialConversationRequest{
		UserID: "60ab81e639ba83faaa1ad7c9",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
	}

	res:= data.GetInitialConversationCRUDOPS(payload) 

	fmt.Println(res)
}

/*func TestGetActionHandlerQuestion(t *testing.T) {

	payload := &data.ActionHandlerRequest{
		ActionHandler:"order",
		UserID: "6083dafb171b889e90c5c7aa",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
	}

	res:= data.GetActionHandlerQuestion(payload) 

	fmt.Println(res)
}*/
