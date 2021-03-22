package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mongo "github.io/zhanchengsong/LocalGuideContentService/database"
	"github.io/zhanchengsong/LocalGuideContentService/model"
)

func HandleGetAllContentRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Contents")
}

func HandleCreateContentRequest(w http.ResponseWriter, r *http.Request) {
	var newContent model.Content
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error handling input data")
		return
	}
	json_err := json.Unmarshal(reqBody, &newContent)
	if json_err != nil {
		fmt.Fprintf(w, "Error parsing json"+json_err.Error())
		return
	}
	mongo_err := mongo.SaveContent(newContent)
	if mongo_err != nil {
		fmt.Fprintf(w, mongo_err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newContent)
}
