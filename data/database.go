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
	}

	var document getData

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetInitialConversationMongo ERROR:")
		log.Error(err)
	}

	if err!=nil{
		isRegistered = false
	} else {
		isRegistered = true
	}

	return isRegistered,err

}