// Package schema provides utilities for encoding and decoding data between the web context and MongoDB BSON format using the Gorilla Schema package.
// It facilitates filtering queries from HTTP requests into BSON-compatible structures.
package schema

import (
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
	gs "github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var decoder = makeDecoder()
var encoder = makeEncoder()

// MakeDecoder creates and configures a new Gorilla schema decoder it sets the decoder to ignore unknown keys that are not present in the destination struct.
func makeDecoder() *gs.Decoder {
	dec := gs.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	return dec
}

// ObjectIdEncoder encodes a MongoDB ObjectID to its hexadecimal string representation for use in query parameters.
func objectIdEncoder(v reflect.Value) string {
	id := v.Interface().(primitive.ObjectID)
	return id.Hex()
}

// MakeEncoder creates and configures a new Gorilla schema encoder and it registers a custom encoder for MongoDB ObjectID types to ensure they are c
// orrectly encoded as hexadecimal strings.
func makeEncoder() *gs.Encoder {
	enc := gs.NewEncoder()
	enc.RegisterEncoder(primitive.ObjectID{}, objectIdEncoder)
	return enc
}

// FilterQuery extracts query parameters from the Gin context, decodes them into a filter structure of type F, and converts them into a
// BSON-compatible format. It returns a bson.M representation of the query parameters or an error if decoding fails.
//
// Parameters:
//   - c: The Gin context containing the HTTP request.
//
// Returns:
//   - A bson.M representation of the query parameters,
//     or an error if decoding fails.
func FilterQuery[F any](c *gin.Context) (bson.M, error) {
	src := c.Request.URL.Query()
	dst := make(url.Values)
	filter := new(F)
	// decode in bson dst | decode the query parameters into the filter structure.
	if err := decoder.Decode(filter, src); err != nil {
		return nil, err
	}
	// Encode the filter structure into BSON-compatible values.
	if err := encoder.Encode(filter, dst); err != nil {
		return nil, err
	}
	query := make(bson.M)
	// Merge the decoded values (dst) into the bson.M.
	for k, v := range src {
		if dst.Has(k) {
			query[k] = v[0]
		}
	}
	return query, nil
}
