package mongo

import (
	"context"
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

//I have used below constants just to hold required database config's.
var (
	CONNECTIONSTRING = os.Getenv("MONGO_CONNECTION")
	DB               = os.Getenv("MONGO_DB")
	CONTENT          = os.Getenv("MONGO_COLLECTION")
)

var credential = options.Credential{
	Username:   os.Getenv("MONGO_USERNAME"),
	Password:   os.Getenv("MONGO_PASSWORD"),
	AuthSource: os.Getenv("MONGO_DB"),
}

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		//Perform connection creation operation only once.
		log.Info(fmt.Sprintf("Connection to Mongo at : %s", CONNECTIONSTRING))
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING).SetAuth(credential)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
			return
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
