package model

type Content struct {
	Titile      string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	VideoId     string `json:"videoId" bson:"videoId"`
}
