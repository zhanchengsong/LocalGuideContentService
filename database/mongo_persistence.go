package mongo

import (
	"context"

	converter "github.io/zhanchengsong/LocalGuideContentService/converters"
	"github.io/zhanchengsong/LocalGuideContentService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDeleteError struct {
	error
	Message string
}

func Error(err MongoDeleteError) string {
	return err.Message
}

func SaveContent(content model.Content) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection(CONTENT)
	_, err = collection.InsertOne(context.TODO(), content)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func GetContentById(contentId string) (model.Content, error) {
	result := model.Content{}
	filter := bson.D{primitive.E{Key: "_id", Value: contentId}}
	client, err := GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(DB).Collection(CONTENT)
	findErr := collection.FindOne(context.TODO(), filter).Decode(&result)
	if findErr != nil {
		return result, findErr
	}
	return result, nil
}

func UpdateContentById(contentId string, content model.Content) error {

	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	updateContent, _ := converter.ConvertContentForUpdate(content)
	collection := client.Database(DB).Collection(CONTENT)
	updater := bson.D{primitive.E{Key: "$set", Value: updateContent}}
	_, u_err := collection.UpdateByID(context.TODO(), contentId, updater)
	if u_err != nil {
		return u_err
	}
	return nil
}

func DeleteContentById(contentId string) (model.Content, error) {
	var deleted = model.Content{}
	client, err := GetMongoClient()
	if err != nil {
		return deleted, err
	}
	collection := client.Database(DB).Collection(CONTENT)

	filter := bson.D{primitive.E{Key: "_id", Value: contentId}}
	findErr := collection.FindOne(context.TODO(), filter).Decode(&deleted)
	if findErr != nil {
		return deleted, findErr
	}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return deleted, err
	}
	if deleteResult.DeletedCount < 1 {
		return deleted, MongoDeleteError{Message: "Cannot delete document"}
	}
	return deleted, nil
}
