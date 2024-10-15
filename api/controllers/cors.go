// Package controllers handles the business logic of the API, including the Preflight function, which handles preflight OPTIONS requests.
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Preflight handles preflight requests by responding with a simple JSON object and returns a 200 OK status with an empty JSON object.
//
// Example:
//
//	curl -X OPTIONS "http://localhost:8080/preflight" \
//	    -H "accept: application/json"
//
// The response will contain an empty JSON object:
// {}
//
// Parameters:
//   - c: *gin.Context - Gin's context to handle the request and response.
//
// Returns:
//   - JSON response with a 200 OK status and an empty object.
func Preflight(c *gin.Context) {
	c.JSON(http.StatusOK, struct{}{})
}
