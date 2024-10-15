package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

// ProfessorRoute initializes the routes related to professor functionality and sets up the "/professor" group and defines the available endpoints.
// This function should be called during the application setup to register the professor-related routes.
//
// The following routes are available:
//
//	OPTIONS /professor:        Calls the Preflight controller to handle CORS preflight requests.
//	GET /professor:            Calls the ProfessorSearch controller to retrieve a list of professors based on search criteria.
//	GET /professor/:id:        Calls the ProfessorById controller to retrieve details of a specific professor by their unique identifier.
//	GET /professor/all:        Calls the ProfessorAll controller to retrieve all professors in the database.
func ProfessorRoute(router *gin.Engine) {
	// All routes related to professors come here
	professorGroup := router.Group("/professor")

	professorGroup.OPTIONS("", controllers.Preflight)
	professorGroup.GET("", controllers.ProfessorSearch)
	professorGroup.GET(":id", controllers.ProfessorById)
	professorGroup.GET("all", controllers.ProfessorAll)
}
