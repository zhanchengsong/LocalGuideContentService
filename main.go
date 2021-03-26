package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	mongo "github.io/zhanchengsong/LocalGuideContentService/database"
	"github.io/zhanchengsong/LocalGuideContentService/handlers"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	// Set up mongoDB connection

	goenverr := godotenv.Load()
	if goenverr != nil {
		log.Error(goenverr.Error())
	}
	_, err := mongo.GetMongoClient()
	if err != nil {
		log.Fatal("Connection to mongodb failed" + err.Error())
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/contents", handlers.HandleGetAllContentRequest).Methods("GET")
	router.HandleFunc("/contents", handlers.HandleCreateContentRequest).Methods("POST")
	router.HandleFunc("/content/{id}", handlers.HandleGetContentByIdRequest).Methods("GET")
	router.HandleFunc("/content/{id}", handlers.HandleUpdateContentRequest).Methods("PATCH")
	router.HandleFunc("/content/{id}", handlers.HandleDeleteContentByIdRequest).Methods("DELETE")
	log.Info(fmt.Sprintf("Service is up and running on port %s", port))

	if err != nil {
		log.Fatal("Port number is not set")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

}
