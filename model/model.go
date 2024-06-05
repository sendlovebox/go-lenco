// Package model holds the object for the request and response payload
package model

import "errors"

const (
	// LogStrRequest log string key
	LogStrRequest = "request"
	// LogStrResponse log string key
	LogStrResponse = "response"
	// BaseURL to for woven finance
	BaseURL = "https://sandbox.lenco.co/access/v1"
	// APIKey for authorization
	APIKey = "bf307db1fcb4749120b81470eff89961367c8760f904d8dffcfa78d633d23e67"
)

var (
	// ErrNetworkError when something goes wrong with the API call
	ErrNetworkError = errors.New("network error")
)

type (
	// Response schema for all endpoints
	Response struct {
		Status    bool        `json:"status"`
		Message   string      `json:"message"`
		ErrorCode *string     `json:"errorCode,omitempty"`
		Errors    interface{} `json:"errors,omitempty"`
		Data      interface{} `json:"data"`
	}
)
