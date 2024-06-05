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
	APIKey = ""
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
