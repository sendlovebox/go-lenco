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
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

//{
//"id": "e0512abf-ecf4-48df-b6ff-40a8bf3b2373",
//"name": "500GB Gold Monthly HyNetFlex Unlimited Plan  (4G Router)",
//"vendor": {
//"id": "70284a19-560d-430d-b1ef-a16fe9107d1f",
//"name": "MTN"
//},
//"amount": {
//"type": "fixed",
//"fixed": "45000.00",
//"minimum": null,
//"maximum": null
//},
//"category": "mobile-data",
//"customerIdLabel": "Mobile Number",
//"commissionPercentage": "0.80"
//},
