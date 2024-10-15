// Package responses provides standardized response structures for API endpoints related to section data.
package responses

import "github.com/UTDNebula/nebula-api/api/schema"

// MultiSectionResponse represents the standardized HTTP response structure for API endpoints that return multiple section data. This response
// includes a status code, a message, and a slice of section data.
type MultiSectionResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []schema.Section `json:"data"`
}

// SingleSectionResponse represents the standardized HTTP response structure for API endpoints that return a single section data. This response
// includes a status code, a message, and a single section data.
type SingleSectionResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    schema.Section `json:"data"`
}
