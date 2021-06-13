package data

import (
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

/*func GetQuestionForActionHandlerMongo (d *ActionHandlerRequest) ([]QuestionObjectArray,error) {
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
			QContext: document.Action.Order.QContext,
			QCustomResponseChoice: document.Action.Order.QCustomResponseChoice,
		}
	}

	if (d.ActionHandler == "feedback") {
		//change feedback to other action handlers
		type HandlerResponseWrapper struct{
			Feedback QuestionAndTypeStruct `json:"feedback"`
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
			Questions: document.Action.Feedback.Questions ,
			QType: document.Action.Feedback.QType ,
			QContext: document.Action.Feedback.QContext,
			QCustomResponseChoice: document.Action.Feedback.QCustomResponseChoice,
		}
	}

	var objArrayResponse []QuestionObjectArray
	
	for i:=0; i < len(resp.Questions); i++ {
		objResponse := QuestionObjectArray{
			Question: resp.Questions[i],
			ResponseType: resp.QType[i],
			QuestionContext: resp.QContext[i],
			CustomResponseChoice: resp.QCustomResponseChoice[i],		
			}
		
			objArrayResponse = append(objArrayResponse,objResponse)
	}

	return objArrayResponse,err
}*/

func GetQuestionForActionHandlerMongo (d *ActionHandlerRequest) ([]QuestionObjectArray,error) {
	collectionName := shashankMongo.DatabaseName.Collection("bot-schema")
	filter := bson.M{"businessid": d.BusinessID}

	//var resp QuestionAndTypeStruct
	var objArrayResponse []QuestionObjectArray
	var err error

	type ActionResponseWrapper struct{
		Action []struct{
			Handler string `json:"handler"`
			Data []struct{
				Questions string `json:"questions"`
				Atype string `json:"atype"`
				Qtype string `json:"qtype"`
				Qcontext string `json:"qcontext"`
				QCard cardObject `json:"qcard"`
				ACard []cardObject `json:"acard"`
				Qcustomresponsechoice []string `json:"qcustomresponsechoice"`	
			} `json:"data"`
		} `json:"action"`
	}

	var document ActionResponseWrapper

	err = collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetQuestionForActionHandlerMongo ERROR:")
		log.Error(err)
	}

	//fmt.Println(document)

	for _,actions:=range document.Action{
		if (actions.Handler == d.ActionHandler){
			for _,data:= range actions.Data{
				objResponse := QuestionObjectArray{
					Question: data.Questions,
					ResponseType: data.Atype,
					QuestionContext: data.Qcontext,
					CustomResponseChoice: data.Qcustomresponsechoice,	
					QuestionType: data.Qtype,
					QestionCard: data.QCard,
					AnswerCard: data.ACard,	
					}
				objArrayResponse = append(objArrayResponse,objResponse)	
			}
		}
	}

	/*if (d.ActionHandler == "feedback") {
		//change feedback to other action handlers
		type HandlerResponseWrapper struct{
			Feedback QuestionAndTypeStruct `json:"feedback"`
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
			Questions: document.Action.Feedback.Questions ,
			QType: document.Action.Feedback.QType ,
			QContext: document.Action.Feedback.QContext,
			QCustomResponseChoice: document.Action.Feedback.QCustomResponseChoice,
		}
	}

	var objArrayResponse []QuestionObjectArray
	
	for i:=0; i < len(resp.Questions); i++ {
		objResponse := QuestionObjectArray{
			Question: resp.Questions[i],
			ResponseType: resp.QType[i],
			QuestionContext: resp.QContext[i],
			CustomResponseChoice: resp.QCustomResponseChoice[i],		
			}
		
			objArrayResponse = append(objArrayResponse,objResponse)
	}*/

	return objArrayResponse,err
}