// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSyncRequest struct {
	SyncUpdate shared.SyncUpdate `request:"mediaType=application/json"`
	// The sync's ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateSyncResponse struct {
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Ok
	UpdateSync200ApplicationJSONAnyOf interface{}
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
