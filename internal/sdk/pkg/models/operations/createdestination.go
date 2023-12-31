// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestination200ApplicationJSONType string

const (
	CreateDestination200ApplicationJSONTypeDestination         CreateDestination200ApplicationJSONType = "Destination"
	CreateDestination200ApplicationJSONTypeValidateErrorJSON   CreateDestination200ApplicationJSONType = "ValidateErrorJSON"
	CreateDestination200ApplicationJSONTypeInternalServerError CreateDestination200ApplicationJSONType = "InternalServerError"
)

type CreateDestination200ApplicationJSON struct {
	Destination         *shared.Destination
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type CreateDestination200ApplicationJSONType
}

func CreateCreateDestination200ApplicationJSONDestination(destination shared.Destination) CreateDestination200ApplicationJSON {
	typ := CreateDestination200ApplicationJSONTypeDestination

	return CreateDestination200ApplicationJSON{
		Destination: &destination,
		Type:        typ,
	}
}

func CreateCreateDestination200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) CreateDestination200ApplicationJSON {
	typ := CreateDestination200ApplicationJSONTypeValidateErrorJSON

	return CreateDestination200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateCreateDestination200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) CreateDestination200ApplicationJSON {
	typ := CreateDestination200ApplicationJSONTypeInternalServerError

	return CreateDestination200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *CreateDestination200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destination := new(shared.Destination)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destination); err == nil {
		u.Destination = destination
		u.Type = CreateDestination200ApplicationJSONTypeDestination
		return nil
	}

	validateErrorJSON := new(shared.ValidateErrorJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&validateErrorJSON); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = CreateDestination200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&internalServerError); err == nil {
		u.InternalServerError = internalServerError
		u.Type = CreateDestination200ApplicationJSONTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateDestination200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.Destination != nil {
		return json.Marshal(u.Destination)
	}

	if u.ValidateErrorJSON != nil {
		return json.Marshal(u.ValidateErrorJSON)
	}

	if u.InternalServerError != nil {
		return json.Marshal(u.InternalServerError)
	}

	return nil, nil
}

type CreateDestinationResponse struct {
	ContentType string
	// Ok
	CreateDestination200ApplicationJSONAnyOf *CreateDestination200ApplicationJSON
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Conflict
	ValidateErrorJSON *shared.ValidateErrorJSON
}
