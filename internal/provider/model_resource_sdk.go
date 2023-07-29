// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"hightouch/internal/sdk/pkg/models/shared"
)

func (r *ModelResourceModel) RefreshFromCreateResponse(resp *shared.ValidateErrorJSON) {
	if r.Details == nil && len(resp.Details) > 0 {
		r.Details = make(map[string]types.String)
		for key, value := range resp.Details {
			result, _ := json.Marshal(value)
			r.Details[key] = types.StringValue(string(result))
		}
	}
	r.Message = types.StringValue(string(resp.Message))
}

func (r *ModelResourceModel) RefreshFromUpdateResponse(resp *shared.ValidateErrorJSON) {
	r.RefreshFromCreateResponse(resp)
}
