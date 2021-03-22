package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mongo "github.io/zhanchengsong/LocalGuideContentService/database"
	"github.io/zhanchengsong/LocalGuideContentService/handlers"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	// Set up mongoDB connection
	_, err := mongo.GetMongoClient()
	if err != nil {
		log.Fatal("Connection to mongodb failed" + err.Error())
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/contents", handlers.HandleGetAllContentRequest).Methods("GET")
	router.HandleFunc("/contents", handlers.HandleCreateContentRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":8443", router))
}
