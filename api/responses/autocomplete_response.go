// Package responses defines the structure of HTTP responses for the API, including the AutocompleteResponse type, which standardizes the response
// format for autocomplete endpoints.
package responses

// AutocompleteResponse represents the standardized HTTP response structure used for endpoints that provide autocomplete functionality. The response
// includes the status code, a message, and any data returned by the query.
//
// Fields:
//
//	Status:  The HTTP status code indicating the success or failure of the request.
//	         For example, 200 for success, 500 for internal server error, etc.
//	Message: A descriptive message corresponding to the status of the request.
//	         For example, "success" or "error".
//	Data:    The data payload returned by the request. This can be any type of data
//	         (interface{}), such as a list of courses or an error message.
type AutocompleteResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
