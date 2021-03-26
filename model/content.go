package model

import "time"

// We use string pointers here to tell if values are not provided (nil) vs empty ("")
type Content struct {
	Id            string    `json:"id,omitempty" bson:"_id,omitempty"`
	Titile        *string   `json:"title,omitempty" bson:"title,omitempty"`
	Description   *string   `json:"description,omitempty" bson:"description,omitempty"`
	VideoId       *string   `json:"videoId,omitempty" bson:"videoId,omitempty"`
	CreatedOn     time.Time `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	LastUpdatedOn time.Time `json:"lastUpdatedOn,omitempty" bson:"lastUpdatedOn,omitempty"`
}

type UpdateContent struct {
	Titile        string    `json:"title,omitempty" bson:"title,omitempty"`
	Description   string    `json:"description,omitempty" bson:"description,omitempty"`
	VideoId       string    `json:"videoId,omitempty" bson:"videoId,omitempty"`
	LastUpdatedOn time.Time `json:"lastUpdatedOn,omitempty" bson:"lastUpdatedOn,omitempty"`
}
