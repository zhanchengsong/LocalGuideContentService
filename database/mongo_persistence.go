package mongo

import (
	"context"

	"github.io/zhanchengsong/LocalGuideContentService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
