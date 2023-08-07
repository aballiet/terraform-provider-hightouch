// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"hightouch/internal/sdk/pkg/models/shared"
	"time"
)

func (r *ModelDataSourceModel) RefreshFromGetResponse(resp *shared.Model) {
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	if r.Custom == nil {
		r.Custom = &ModelCreateCustom{}
	}
	if resp.Custom == nil {
		r.Custom = nil
	} else {
		r.Custom = &ModelCreateCustom{}
		queryResult, _ := json.Marshal(resp.Custom.Query)
		r.Custom.Query = types.StringValue(string(queryResult))
	}
	if r.Dbt == nil {
		r.Dbt = &ModelCreateDbt{}
	}
	if resp.Dbt == nil {
		r.Dbt = nil
	} else {
		r.Dbt = &ModelCreateDbt{}
		r.Dbt.CompiledSQL = types.StringValue(resp.Dbt.CompiledSQL)
		r.Dbt.Database = types.StringValue(resp.Dbt.Database)
		r.Dbt.DbtUniqueID = types.StringValue(resp.Dbt.DbtUniqueID)
		r.Dbt.ModelID = types.StringValue(resp.Dbt.ModelID)
		r.Dbt.Name = types.StringValue(resp.Dbt.Name)
		r.Dbt.RawSQL = types.StringValue(resp.Dbt.RawSQL)
		r.Dbt.Schema = types.StringValue(resp.Dbt.Schema)
	}
	r.ID = types.StringValue(resp.ID)
	r.IsSchema = types.BoolValue(resp.IsSchema)
	r.Name = types.StringValue(resp.Name)
	r.PrimaryKey = types.StringValue(resp.PrimaryKey)
	r.QueryType = types.StringValue(resp.QueryType)
	if r.Raw == nil {
		r.Raw = &ModelCreateRaw{}
	}
	if resp.Raw == nil {
		r.Raw = nil
	} else {
		r.Raw = &ModelCreateRaw{}
		r.Raw.SQL = types.StringValue(resp.Raw.SQL)
	}
	r.Slug = types.StringValue(resp.Slug)
	r.SourceID = types.StringValue(resp.SourceID)
	r.Syncs = nil
	for _, v := range resp.Syncs {
		r.Syncs = append(r.Syncs, types.StringValue(v))
	}
	if r.Table == nil {
		r.Table = &ModelCreateTable{}
	}
	if resp.Table == nil {
		r.Table = nil
	} else {
		r.Table = &ModelCreateTable{}
		r.Table.Name = types.StringValue(resp.Table.Name)
	}
	if r.Tags == nil && len(resp.Tags) > 0 {
		r.Tags = make(map[string]types.String)
		for key, value := range resp.Tags {
			r.Tags[key] = types.StringValue(value)
		}
	}
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	if r.Visual == nil {
		r.Visual = &ModelCreateVisual{}
	}
	if resp.Visual == nil {
		r.Visual = nil
	} else {
		r.Visual = &ModelCreateVisual{}
		filterResult, _ := json.Marshal(resp.Visual.Filter)
		r.Visual.Filter = types.StringValue(string(filterResult))
		r.Visual.ParentID = types.StringValue(resp.Visual.ParentID)
		r.Visual.PrimaryLabel = types.StringValue(resp.Visual.PrimaryLabel)
		r.Visual.SecondaryLabel = types.StringValue(resp.Visual.SecondaryLabel)
	}
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}