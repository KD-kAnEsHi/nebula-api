package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

// GradesRoute initializes the routes related to grade functionality and sets up the "/grades" group and defines the available endpoints.
// This function should be called during the application setup to register the grade-related routes.
//
// The following routes are available:
//
//	OPTIONS /grades:       Calls the Preflight controller to handle CORS preflight requests.
//	GET /grades/semester:  Calls the GradeAggregationSemester controller  to retrieve aggregated grades by semester.
//	GET /grades/overall:    Calls the GradesAggregationOverall controller to retrieve overall grade aggregations.
func GradesRoute(router *gin.Engine) {
	// All routes related to sections come here
	gradesGroup := router.Group("/grades")

	gradesGroup.OPTIONS("", controllers.Preflight)

	// @TODO: Do we need this?
	// ---- gradesGroup.OPTIONS("semester", controllers.Preflight)
	// ---- gradesGroup.OPTIONS("overall", controllers.Preflight)

	gradesGroup.GET("semester", controllers.GradeAggregationSemester())
	gradesGroup.GET("overall", controllers.GradesAggregationOverall())
}
