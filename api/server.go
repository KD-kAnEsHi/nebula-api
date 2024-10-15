// Package main is the entry point for the Nebula Labs API application and it sets up the API server, configures routes, middleware, and handles requests.
package main

import (
	"github.com/UTDNebula/nebula-api/api/common/log"
	"github.com/UTDNebula/nebula-api/api/configs"
	_ "github.com/UTDNebula/nebula-api/api/docs"
	"github.com/UTDNebula/nebula-api/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title nebula-api
// @description The public Nebula Labs API for access to pertinent UT Dallas data
// @version 0.1.0
// @host nebula-api-2lntm5dxoflqn.apigateway.nebula-api-368223.cloud.goog
// @schemes http
// @x-google-backend {"address": "REDACTED"}
// @x-google-endpoints [{"name": "nebula-api-2lntm5dxoflqn.apigateway.nebula-api-368223.cloud.goog", "allowCors": true}]
// @x-google-management {"metrics": [{"name": "read-requests", "displayName": "Read Requests CUSTOM", "valueType": "INT64", "metricKind": "DELTA"}], "quota": {"limits": [{"name": "read-limit", "metric": "read-requests", "unit": "1/min/{project}", "values": {"STANDARD": 1000}}]}}
// @securitydefinitions.apikey apiKey
// @name x-api-key
// @in header

// Main initializes the application, sets up the router, and starts the API server and connects to the database, configures middleware, and registers
// routes for the application.
//
// Middleware functions include:
//   - CORS: Enables Cross-Origin Resource Sharing.
//   - LogRequest: Logs incoming requests for monitoring purposes.
//
// Swagger documentation is hosted at the "/swagger/*any" endpoint.
//
// Example usage:
//
//	go run main.go
func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Establish the connection to the database
	configs.ConnectDB()

	// Configure Gin Router
	router := gin.New()

	// Enable CORS
	router.Use(CORS())

	// Enable Logging
	router.Use(LogRequest)

	// Setup swagger-ui hosted
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Connect Routes
	routes.CourseRoute(router)
	routes.SectionRoute(router)
	routes.ProfessorRoute(router)
	routes.GradesRoute(router)
	routes.AutocompleteRoute(router)
	routes.StorageRoute(router)

	// Retrieve the port string to serve traffic on
	portString := configs.GetPortString()

	// Serve Traffic
	router.Run(portString)
	log.Logger.Debug().Str("port", portString).Msg("Listening to port")
}

// CORS returns a gin.HandlerFunc that sets up the CORS headers for incoming requests and allows requests from any origin and specifies which
// headers and methods are permitted. For preflight requests (OPTIONS), it responds with a 204 No Content status.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, x-api-key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.IndentedJSON(204, "")
			return
		}

		c.Next()
	}
}

// This function logs tdetails about incoming HTTP requests, including the method, path, and host, for monitoring purposes.
func LogRequest(c *gin.Context) {
	log.Logger.Info().
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("host", c.Request.Host).
		Send()
	c.Next()
}
