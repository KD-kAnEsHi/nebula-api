package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

// SectionRoute initializes the routes related to section functionality and sets up the "/section" group and defines the available endpoints.
// This function should be called during the application setup to register the section-related routes.
//
// The following routes are available:
//
//	OPTIONS /section:               Calls the Preflight controller to handle CORS preflight requests.
//	GET /section:                   Calls the SectionSearch controller to retrieve a list of sections based on search criteria.
//	GET /section/:id:               Calls the SectionById controller to retrieve details of a specific section by its unique identifier.
//	GET /section/:id/evaluation:    Calls the EvalBySectionID controller to retrieve evaluations related to a specific section.
func SectionRoute(router *gin.Engine) {
	// All routes related to sections come here
	sectionGroup := router.Group("/section")

	sectionGroup.OPTIONS("", controllers.Preflight)
	sectionGroup.GET("", controllers.SectionSearch)
	sectionGroup.GET(":id", controllers.SectionById)
	sectionGroup.GET(":id/evaluation", controllers.EvalBySectionID)
}
