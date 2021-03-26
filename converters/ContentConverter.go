package converter

import (
	"time"

	"github.io/zhanchengsong/LocalGuideContentService/model"
)

func ConvertContentForUpdate(content model.Content) (model.UpdateContent, error) {
	var result = model.UpdateContent{}
	if content.Titile != nil {
		result.Titile = *content.Titile
	}
	if content.VideoId != nil {
		result.VideoId = *content.VideoId
	}
	if content.Description != nil {
		result.Description = *content.Description
	}
	result.LastUpdatedOn = time.Now()
	return result, nil
}
