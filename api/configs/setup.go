// Package configs handles database connections and query options for interacting with MongoDB.
// It provides functions for establishing a connection to MongoDB, retrieving collections,
// and creating query options such as limit and offset for data retrieval.
package configs

import (
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/UTDNebula/nebula-api/api/common/log"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBSingleton represents a singleton for managing a MongoDB client connection. It ensures that
// only one instance of the client is created and shared across the application.
type DBSingleton struct {
	client *mongo.Client
}

var dbInstance *DBSingleton
var once sync.Once

// ConnectDB establishes a connection to the MongoDB database using the URI retrieved from environment variables.
// It make sure that the connection is only created once (singleton pattern) and pings the database to verify the connection.
// The function then returns a MongoDB client that can be used to interact with the database.
//
// Example usage:
//
//	client := configs.ConnectDB()
//	collection := client.Database("combinedDB").Collection("users")
func ConnectDB() *mongo.Client {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// Retrieve the MongoDB URI from the environment
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(GetEnvMongoURI()))
		if err != nil {
			log.WriteErrorMsg("Unable to create MongoDB client")
			os.Exit(1) // Terminate if program is unable to create MongoDB client
		}

		defer cancel()

		// Ping the database to check connectivity
		err = client.Ping(ctx, nil)
		if err != nil {
			log.WriteErrorMsg("Unable to ping database")
			os.Exit(1) // Terminate if program is unable to ping to the database
		}

		log.WriteDebug("Connected to MongoDB")

		// Store the client in the singleton instance
		dbInstance = &DBSingleton{
			client: client,
		}

	})

	return dbInstance.client
}

// GetCollection retrieves a MongoDB collection from the database named "combinedDB".
// It ensures the database connection is established before retrieving the collection.
//
// Example usage:
//
//	collection := configs.GetCollection("users")
//	filter := bson.M{"name": "John Doe"}
//	result := collection.FindOne(context.Background(), filter)
//
// Params:
//
//	collectionName - The name of the collection to retrieve.
//
// Returns a MongoDB collection object for performing operations like Find, Insert, etc..
func GetCollection(collectionName string) *mongo.Collection {
	client := ConnectDB()
	collection := client.Database("combinedDB").Collection(collectionName)
	return collection
}

// GetOptionLimit generates a MongoDB FindOptions object with a limit and offset for paginated queries.
// It retrieves the 'offset' query parameter from the request context and applies it along with a limit
// from the environment variables or a default value.
//
// Example usage:
//
//	var query bson.M
//	findOptions, err := configs.GetOptionLimit(&query, c)
//	if err != nil {
//	  log.Println("Error applying limit/offset:", err)
//	}
//
// Params:
//
//	query - A pointer to a BSON query object, where "offset" is removed (if present).
//	c     - A Gin context object to extract query parameters.
//
// Returns a pointer to MongoDB FindOptions ( *options.FindOptions )with the applied limit. Returns error if any
func GetOptionLimit(query *bson.M, c *gin.Context) (*options.FindOptions, error) {
	delete(*query, "offset") // removes offset (if present) in query --offset is not field in collections
	// Default limit from environment variables
	var limit int64 = GetEnvLimit()

	// parses offset if included in the query
	var offset int64
	var err error

	if c.Query("offset") == "" {
		offset = 0 // Default value if no offset is specified
	} else {
		offset, err = strconv.ParseInt(c.Query("offset"), 10, 64)
		if err != nil {
			// If parsing fails, use default offset and return the error
			return options.Find().SetSkip(0).SetLimit(limit), err // default value for offset
		}
	}
	// Return the FindOptions with the applied offset and limit
	return options.Find().SetSkip(offset).SetLimit(limit), err
}
