// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSourceRequest struct {
	SourceUpdate shared.SourceUpdate `request:"mediaType=application/json"`
	// The source's ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateSourceResponse struct {
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Ok
	UpdateSource200ApplicationJSONAnyOf interface{}
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
