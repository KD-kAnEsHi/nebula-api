// Package responses provides standardized response structures for API endpoints related to course data and evaluations.
package responses

// GradeResponse represents the standardized HTTP response structure for API endpoints that return grade data. This response includes
// a status code, a message, and the grade data itself.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 200 for success).
//	Message: A brief description of the result of the request (e.g., "success" or "error").
//	Data:    The grade data, which can be of any type, represented as an empty interface.
type GradeResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
