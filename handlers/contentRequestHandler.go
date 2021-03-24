package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	mongo "github.io/zhanchengsong/LocalGuideContentService/database"
	"github.io/zhanchengsong/LocalGuideContentService/model"
)

type requestError struct {
	Message string
}

func HandleGetAllContentRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	json.NewEncoder(w).Encode(requestError{Message: "Not implemented"})
}

func HandleGetContentByIdRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	contentId, ok := mux.Vars(r)["id"]
	log.WithFields(log.Fields{
		"handler": "HandleGetContentByIdRequest",
	}).Debug(contentId)
	if !ok || len(contentId) < 1 {
		log.WithFields(log.Fields{
			"handler": "HandleGetContentByIdRequest",
		}).Error("Missing contentId in query params")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestError{Message: "Missing contentId"})
		return
	}

	contentResult, err := mongo.GetContentById(contentId)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "HandleGetContentByIdRequest",
		}).Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(requestError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(contentResult)
	elapsed := time.Since(start).Milliseconds()
	log.WithFields(log.Fields{
		"handler": "HandleGetContentByIdRequest",
	}).Info("Request handled in " + fmt.Sprintf("%d ms", elapsed))

}

func HandleCreateContentRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var newContent model.Content
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "HandleCreateContentRequest",
		}).Error("Cannot read request body: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestError{Message: err.Error()})
		return
	}
	json_err := json.Unmarshal(reqBody, &newContent)
	if json_err != nil {
		log.WithFields(log.Fields{
			"handler": "HandleCreateContentRequest",
		}).Error("Cannot parse request body: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestError{Message: json_err.Error()})
		return
	}
	// Fill in timestamps and id
	newContent.CreatedOn = time.Now()
	newContent.LastUpdatedOn = time.Now()
	newContent.Id = uuid.NewString()
	mongo_err := mongo.SaveContent(newContent)
	if mongo_err != nil {
		log.WithFields(log.Fields{
			"handler": "HandleCreateContentRequest",
		}).Error("MongoDB error: " + mongo_err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(requestError{Message: mongo_err.Error()})
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newContent)
	elapsed := time.Since(start).Milliseconds()
	log.WithFields(log.Fields{
		"handler": "HandleCreateContentRequest",
	}).Info("Request handled in " + fmt.Sprintf("%d ms", elapsed))
}
