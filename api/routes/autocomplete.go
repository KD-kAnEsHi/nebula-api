// Package routes defines the API routes for the Nebula application.
package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

// AutocompleteRoute initializes the routes related to autocomplete functionality. It sets up the "/autocomplete" group and defines the available endpoints.
// This function should be called during the application setup to register the autocomplete-related routes.
//
// The following routes are available:
//
//	GET /autocomplete/dag:  Calls the AutocompleteDAG controller to handle autocomplete requests for Directed Acyclic Graphs (DAG).
func AutocompleteRoute(router *gin.Engine) {
	// All routes related to autocomplete come here
	autocompleteGroup := router.Group("/autocomplete")

	autocompleteGroup.GET("/dag", controllers.AutocompleteDAG)
}
