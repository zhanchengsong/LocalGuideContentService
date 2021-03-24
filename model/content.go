package model

import "time"

type Content struct {
	Id            string    `json:"id" bson:"_id"`
	Titile        string    `json:"title" bson:"title"`
	Description   string    `json:"description" bson:"description"`
	VideoId       string    `json:"videoId" bson:"videoId"`
	CreatedOn     time.Time `json:"createdOn" bson:"createdOn"`
	LastUpdatedOn time.Time `json:"lastUpdatedOn" bson:"lastUpdatedOn"`
}
