// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunCustom200ApplicationJSONType string

const (
	TriggerRunCustom200ApplicationJSONTypeTriggerRunOutput  TriggerRunCustom200ApplicationJSONType = "TriggerRunOutput"
	TriggerRunCustom200ApplicationJSONTypeValidateErrorJSON TriggerRunCustom200ApplicationJSONType = "ValidateErrorJSON"
)

type TriggerRunCustom200ApplicationJSON struct {
	TriggerRunOutput  *shared.TriggerRunOutput
	ValidateErrorJSON *shared.ValidateErrorJSON

	Type TriggerRunCustom200ApplicationJSONType
}

func CreateTriggerRunCustom200ApplicationJSONTriggerRunOutput(triggerRunOutput shared.TriggerRunOutput) TriggerRunCustom200ApplicationJSON {
	typ := TriggerRunCustom200ApplicationJSONTypeTriggerRunOutput

	return TriggerRunCustom200ApplicationJSON{
		TriggerRunOutput: &triggerRunOutput,
		Type:             typ,
	}
}

func CreateTriggerRunCustom200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) TriggerRunCustom200ApplicationJSON {
	typ := TriggerRunCustom200ApplicationJSONTypeValidateErrorJSON

	return TriggerRunCustom200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func (u *TriggerRunCustom200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	triggerRunOutput := new(shared.TriggerRunOutput)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&triggerRunOutput); err == nil {
		u.TriggerRunOutput = triggerRunOutput
		u.Type = TriggerRunCustom200ApplicationJSONTypeTriggerRunOutput
		return nil
	}

	validateErrorJSON := new(shared.ValidateErrorJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&validateErrorJSON); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = TriggerRunCustom200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TriggerRunCustom200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.TriggerRunOutput != nil {
		return json.Marshal(u.TriggerRunOutput)
	}

	if u.ValidateErrorJSON != nil {
		return json.Marshal(u.ValidateErrorJSON)
	}

	return nil, nil
}

type TriggerRunCustomResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Ok
	TriggerRunCustom200ApplicationJSONAnyOf *TriggerRunCustom200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
