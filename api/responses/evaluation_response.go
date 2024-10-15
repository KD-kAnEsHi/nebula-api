// Package responses provides standardized response structures for API endpoints related to course data and evaluations.
package responses

import "github.com/UTDNebula/nebula-api/api/schema"

// EvaluationResponse represents the standardized HTTP response structure for API endpoints that return evaluation data. This response includes
// a status code, a message, and the evaluation data itself.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 200 for success).
//	Message: A brief description of the result of the request (e.g., "success" or "error").
//	Data:    An instance of schema.Evaluation containing the evaluation details.
type EvaluationResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    schema.Evaluation `json:"data"`
}
