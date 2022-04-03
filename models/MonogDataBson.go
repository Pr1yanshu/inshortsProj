package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoDocStruct struct {
	Id   primitive.ObjectID    `json:"id" bson:"_id"`
	Data map[string]RegionData `json:"data" bson:"data"`
	Name string                `json:"name" bson:"name"`
}
