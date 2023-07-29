// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

// ListSourceOrderBy - specify the order
type ListSourceOrderBy string

const (
	ListSourceOrderByID        ListSourceOrderBy = "id"
	ListSourceOrderByName      ListSourceOrderBy = "name"
	ListSourceOrderBySlug      ListSourceOrderBy = "slug"
	ListSourceOrderByCreatedAt ListSourceOrderBy = "createdAt"
	ListSourceOrderByUpdatedAt ListSourceOrderBy = "updatedAt"
)

func (e ListSourceOrderBy) ToPointer() *ListSourceOrderBy {
	return &e
}

func (e *ListSourceOrderBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "name":
		fallthrough
	case "slug":
		fallthrough
	case "createdAt":
		fallthrough
	case "updatedAt":
		*e = ListSourceOrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSourceOrderBy: %v", v)
	}
}

type ListSourceRequest struct {
	// limit the number of objects returned (default is 100)
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// filter based on name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// set the offset on results (for pagination)
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *ListSourceOrderBy `queryParam:"style=form,explode=true,name=orderBy"`
	// filter based on slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

// ListSource200ApplicationJSON - Ok
type ListSource200ApplicationJSON struct {
	Data []shared.Source `json:"data"`
}

type ListSourceResponse struct {
	ContentType string
	// Ok
	ListSource200ApplicationJSONObject *ListSource200ApplicationJSON
	StatusCode                         int
	RawResponse                        *http.Response
}