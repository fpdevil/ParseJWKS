package utils

import (
	"encoding/pem"
	"net/http"
)

// JWK represents the model fields of a standard JWK key set
// exposed by a provider as per the standards and at a minimum
// the provider should at least expose those values
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	E   string `json:"e"`
	N   string `json:"n"`
}

// JWKS represents a set of JWK units if more than one are made
// available by the provider
type JWKS struct {
	Keys []JWK `json:"keys"`
}

// Result struct holds the http response and probable errors as
// obtained from making a HTTP call
type Result struct {
	Response *http.Response // http response
	Error    error          // error response
}

// Data struct represents a collective unit of various fields
// that are being passed to a function for parsing the http response
type Data struct {
	Url      string
	Status   string
	Response []byte
	Error    error
	Count    int
}

// KeyData struct represents the data fields that would be provided
// by the function that parses the JWKS information
type KeyData struct {
	Blocks map[string]pem.Block
	Error  error
}
