// Package responses provides standardized response structures for API endpoints related to course data. It includes responses for both multiple
// and single course retrieval.
package responses

import "github.com/UTDNebula/nebula-api/api/schema"

// MultiCourseResponse represents the standardized HTTP response structure for API endpoints that return multiple courses. This response includes
// a status code, a message, and an array of course data.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 200 for success).
//	Message: A brief description of the result of the request (e.g., "success" or "error").
//	Data:    A slice of Course objects containing the details of the courses.
type MultiCourseResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []schema.Course `json:"data"`
}

// SingleCourseResponse represents the standardized HTTP response structure for API endpoints that return a single course. This response
// includes a status code, a message, and a single course object.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 200 for success).
//	Message: A brief description of the result of the request (e.g., "success" or "error").
//	Data:    A single Course object containing the details of the course.
type SingleCourseResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    schema.Course `json:"data"`
}
