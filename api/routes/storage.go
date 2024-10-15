package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

// StorageRoute initializes the routes related to storage functionality and sets up the "/storage" group and defines the available endpoints.
// This function should be called during the application setup to register the storage-related routes.
//
// The following routes are available (currently commented out):
//
//	OPTIONS /storage:                		Calls the Preflight controller to handle CORS preflight requests.
//	GET /storage/:bucket:           		Calls the BucketInfo controller to retrieve information about a specific storage bucket.
//	GET /storage/:bucket/info/:objectID: 	Calls the ObjectInfo controller to retrieve information about a specific object within the given bucket.
//	POST /storage/:bucket/post/:objectID: 	Calls the PostObject controller to upload an object to the specified bucket.
//	GET /storage/:bucket/get/:objectID: 	Calls the GetObject controller to retrieve a specific object from the bucket.
func StorageRoute(router *gin.Engine) {
	// All routes related to storage come here
	storageGroup := router.Group("/storage")

	storageGroup.OPTIONS("", controllers.Preflight)
	/*
		storageGroup.GET(":bucket", controllers.BucketInfo())
		storageGroup.GET(":bucket/info/:objectID", controllers.ObjectInfo())
		storageGroup.POST(":bucket/post/:objectID", controllers.PostObject())
		storageGroup.GET(":bucket/get/:objectID", controllers.GetObject())
	*/
}
