// Package responses provides standardized response structures for API endpoints related to professor data.
package responses

import "github.com/UTDNebula/nebula-api/api/schema"

// MultiProfessorResponse represents the standardized HTTP response structure for API endpoints that return multiple professor data. This
// response includes a status code, a message, and a slice of professor data.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 200 for success).
//	Message: A brief description of the result of the request (e.g., "success" or "error").
//	Data:    A slice of schema.Professor containing information about the professors.
type MultiProfessorResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    []schema.Professor `json:"data"`
}

// SingleProfessorResponse represents the standardized HTTP response structure for API endpoints that return a single professor data. This response
// includes a status code, a message, and a single professor data.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 200 for success).
//	Message: A brief description of the result of the request (e.g., "success" or "error").
//	Data:    The schema.Professor containing information about the professor.
type SingleProfessorResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    schema.Professor `json:"data"`
}
