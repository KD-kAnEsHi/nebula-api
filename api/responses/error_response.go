// Package responses provides standardized response structures for API endpoints related to course data. It includes responses for errors that
// may occur during API requests.
package responses

// ErrorResponse represents the standardized HTTP response structure for API endpoints when an error occurs. This response includes a status
// code, a message, and additional error details.
//
// Fields:
//
//	Status:  The HTTP status code indicating the result of the request (e.g., 400 for bad request).
//	Message: A brief description of the result of the request (e.g., "error" or "not found").
//	Data:    A string providing more specific error details or information.
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"error"`
}
