package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

// CourseRoute initializes the routes related to course functionality and sets up the "/course" group and defines the available endpoints.
// This function should be called during the application setup to register the course-related routes.
//
// The following routes are available:
//
//	OPTIONS /course:        Calls the Preflight controller to handle CORS preflight requests.
//	GET /course:           Calls the CourseSearch controller to search for courses based on provided query parameters.
//	GET /course/:id:       Calls the CourseById controller to retrieve a course by its ID.
//	GET /course/all:       Calls the CourseAll controller to retrieve all available courses.
func CourseRoute(router *gin.Engine) {
	// All routes related to courses come here
	courseGroup := router.Group("/course")

	courseGroup.OPTIONS("", controllers.Preflight)
	courseGroup.GET("", controllers.CourseSearch)
	courseGroup.GET(":id", controllers.CourseById)
	courseGroup.GET("all", controllers.CourseAll)
}
