// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSyncRequest struct {
	// The id of the sync
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetSyncResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Ok
	Sync *shared.Sync
}
