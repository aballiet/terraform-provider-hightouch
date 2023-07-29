// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationResponse struct {
	ContentType string
	// Ok
	CreateDestination200ApplicationJSONAnyOf interface{}
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Conflict
	ValidateErrorJSON *shared.ValidateErrorJSON
}
