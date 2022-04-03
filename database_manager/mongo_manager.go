package database_manager

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"inshortsProj/constant"
	"inshortsProj/logger"
	"inshortsProj/models"
	"log"
	"runtime/debug"
)

var (
	client *mongo.Client
	Ctx    = context.TODO()
)

func init() {
	initializeMongoDB()
}

/*Setup opens a database connection to mongodb*/
func initializeMongoDB() {
	connectionURI := "mongodb+srv://priyanshu:priyanshu1234@cluster0.o2vox.mongodb.net/test/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	var err error
	client, err = mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetStateCollectionData(ctx *gin.Context) (models.MongoDocStruct, error) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic in getStateCollectionData ,General Error: " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			fmt.Println(ErrorString)
			panic(ErrorString)
		}
	}()
	if client == nil {
		initializeMongoDB()
	}
	db := client.Database("CovidDatabase")
	stateCollection := db.Collection("StateData")

	var document models.MongoDocStruct
	cursor, err := stateCollection.Find(ctx, bson.D{})
	if err != nil {
		defer cursor.Close(ctx)
		return document, err
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&document)
		if err != nil {
			return document, err
		}
		break
	}

	return document, nil

}

func SetStateCollectionData(ctx *gin.Context, response map[string]models.RegionData) error {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic in setStateCollectionData ,General Error: " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			fmt.Println(ErrorString)
			logger.LogErrorForScalyr(ErrorString, "SetStateCollectionData", "api/updateCovidData", "")
			panic(ErrorString)
		}
	}()
	if client == nil {
		initializeMongoDB()
	}

	var finalDoc models.MongoDocStruct
	finalDoc.Data = response
	finalDoc.Name = constant.MONGO_DOC_NAME
	db := client.Database("CovidDatabase")
	stateCollection := db.Collection("StateData")
	// purge old data
	_, err := stateCollection.DeleteMany(ctx, bson.M{"name": constant.MONGO_DOC_NAME})
	if err != nil {
		fmt.Println("error in purging old data in db : " + err.Error())
		return err
	}
	// insert latest data
	_, err = stateCollection.InsertOne(Ctx, finalDoc)
	if err != nil {
		fmt.Println("error in inserting to database : " + err.Error())
		return err
	}
	return nil
}
