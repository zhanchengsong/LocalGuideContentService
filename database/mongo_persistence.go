package mongo

import (
	"context"

	"github.io/zhanchengsong/LocalGuideContentService/model"
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
