package hellomongo

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/mongodb"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
	expConf "github.com/jabong/florest-core/src/examples/config"
	"reflect"
)

type mongoNode struct {
	id string
}

func (a *mongoNode) SetID(id string) {
	a.id = id
}

func (a mongoNode) GetID() (id string, err error) {
	return a.id, nil
}

func (a mongoNode) Name() string {
	return "mongoNode"
}

// an example for mongo document
type employeeInfo struct {
	ID   string
	Type string
}

func (a mongoNode) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	appConfig, err := expConf.GetExampleAppConfig()
	if err != nil {
		msg := "Mongo App Config Not Correct"
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	mongoConf := appConfig.Mongo
	if mongoConf == nil {
		msg := " Mongo Config is Missing"
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	collection := "florest"

	err = mongodb.Set("mymongo", mongoConf, reflect.TypeOf(mongodb.MongoDriver{}))
	if err != nil {
		msg := fmt.Sprintf("Mongo Set has error %v", err)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	db, errG := mongodb.Get("mymongo")
	if errG != nil {
		msg := fmt.Sprintf("Mongo Config Not Correct %v", errG)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}
	var query map[string]interface{}

	// insert
	if errI := db.Insert(collection, &employeeInfo{ID: "123", Type: "Manager"}); errI != nil {
		msg := fmt.Sprintf("Insert to Mongo Failed - %v", errI)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	// update
	query = make(map[string]interface{}, 1)
	query["id"] = "123"
	if errU := db.Update(collection, query, &employeeInfo{ID: "123", Type: "Director"}); errU != nil {
		msg := fmt.Sprintf("Mongo Update Failed %v", errU)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: fmt.Sprintf("Mongo Update Failed %v", errU)}
	}

	// find one
	if _, errF := db.FindOne(collection, query); errF != nil {
		msg := fmt.Sprintf("Find from Mongo Failed - %v", errF)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	// find all
	if _, errF := db.FindAll(collection, query); errF != nil {
		msg := fmt.Sprintf("Find All from Mongo Failed - %v", errF)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	// remove
	if errR := db.Remove(collection, query); errR != nil {
		msg := fmt.Sprintf("Remove from Mongo Failed - %v", errR)
		logger.Error(msg)
		return io, &constants.AppError{Code: constants.InvalidErrorCode, Message: msg}
	}

	//Business Logic
	io.IOData.Set(constants.Result, "Insert, Update, Find, FindAll, Remove operation sucessful on Mongo employee collection")
	return io, nil
}
