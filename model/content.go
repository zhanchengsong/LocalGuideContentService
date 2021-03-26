package model

import "time"

type Content struct {
	Id            string    `json:"id,omitempty" bson:"_id"`
	Titile        string    `json:"title,omitempty" bson:"title"`
	Description   string    `json:"description,omitempty" bson:"description"`
	VideoId       string    `json:"videoId,omitempty" bson:"videoId"`
	CreatedOn     time.Time `json:"createdOn,omitempty" bson:"createdOn"`
	LastUpdatedOn time.Time `json:"lastUpdatedOn,omitempty" bson:"lastUpdatedOn"`
}
