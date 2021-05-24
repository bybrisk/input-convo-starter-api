package data

import (
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func GetInitialConversationMongo(businessId string) (*InitialConversationResponse,error){
	collectionName := shashankMongo.DatabaseName.Collection("bot-schema")
	filter := bson.M{"businessid": businessId}

	type InitialConversationResponseWrapper struct{
		Initialise InitialConversationResponse `json:"initialise"`
	}

	var document InitialConversationResponseWrapper

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetInitialConversationMongo ERROR:")
		log.Error(err)
	}

	return &document.Initialise,err
}

func IsUserRegistered(docID string) (bool,error) {
	var isRegistered bool

	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}

	type getData struct {
		Phonenumber string `json:"phonenumber"`
		UserName string `json:"username"`
	}

	var document getData

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetInitialConversationMongo ERROR:")
		log.Error(err)
	}

	if err!=nil {
		isRegistered = false
	} else {
		if document.UserName == "" {
			isRegistered = false
		} else {
			isRegistered = true
		}
	}
	return isRegistered,err
}

func GetQuestionForActionHandlerMongo (d *ActionHandlerRequest) (*QuestionAndTypeStruct,error) {
	collectionName := shashankMongo.DatabaseName.Collection("bot-schema")
	filter := bson.M{"businessid": d.BusinessID}

	var resp QuestionAndTypeStruct
	var err error

	//ActionHandler specific struct
	if (d.ActionHandler == "order") {
		//change Order to other action handlers
		type HandlerResponseWrapper struct{
			Order QuestionAndTypeStruct `json:"order"`
		}
		type ActionResponseWrapper struct{
			Action HandlerResponseWrapper `json:"action"`
		}
	
		var document ActionResponseWrapper

		err = collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
		if err != nil {
			log.Error("GetQuestionForActionHandlerMongo ERROR:")
			log.Error(err)
		}

		resp = QuestionAndTypeStruct{
			Questions: document.Action.Order.Questions ,
			QType: document.Action.Order.QType ,
		}
	}

	return &resp,err
}