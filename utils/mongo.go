package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type mongoUtils struct {
}

type MongoUtils interface {
	ConvertToObjectID(id string) (interface{}, error)
}

func (u *mongoUtils) ConvertToObjectID(id string) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	return objID, err
}

func NewMongoUtils() MongoUtils {
	return &mongoUtils{}
}
