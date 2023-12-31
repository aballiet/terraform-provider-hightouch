// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"hightouch/internal/sdk/pkg/models/shared"
	"time"
)

func (r *ModelResourceModel) ToCreateSDKType() *shared.ModelCreate {
	var custom *shared.ModelCreateCustom
	if r.Custom != nil {
		var query interface{}
		_ = json.Unmarshal([]byte(r.Custom.Query.ValueString()), &query)
		custom = &shared.ModelCreateCustom{
			Query: query,
		}
	}
	var dbt *shared.ModelCreateDbt
	if r.Dbt != nil {
		modelID := r.Dbt.ModelID.ValueString()
		dbt = &shared.ModelCreateDbt{
			ModelID: modelID,
		}
	}
	isSchema := r.IsSchema.ValueBool()
	name := r.Name.ValueString()
	primaryKey := r.PrimaryKey.ValueString()
	queryType := r.QueryType.ValueString()
	var raw *shared.ModelCreateRaw
	if r.Raw != nil {
		sql := r.Raw.SQL.ValueString()
		raw = &shared.ModelCreateRaw{
			SQL: sql,
		}
	}
	slug := r.Slug.ValueString()
	sourceID := r.SourceID.ValueString()
	var table *shared.ModelCreateTable
	if r.Table != nil {
		name1 := r.Table.Name.ValueString()
		table = &shared.ModelCreateTable{
			Name: name1,
		}
	}
	var visual *shared.ModelCreateVisual
	if r.Visual != nil {
		var filter interface{}
		_ = json.Unmarshal([]byte(r.Visual.Filter.ValueString()), &filter)
		parentID := r.Visual.ParentID.ValueString()
		primaryLabel := r.Visual.PrimaryLabel.ValueString()
		secondaryLabel := r.Visual.SecondaryLabel.ValueString()
		visual = &shared.ModelCreateVisual{
			Filter:         filter,
			ParentID:       parentID,
			PrimaryLabel:   primaryLabel,
			SecondaryLabel: secondaryLabel,
		}
	}
	out := shared.ModelCreate{
		Custom:     custom,
		Dbt:        dbt,
		IsSchema:   isSchema,
		Name:       name,
		PrimaryKey: primaryKey,
		QueryType:  queryType,
		Raw:        raw,
		Slug:       slug,
		SourceID:   sourceID,
		Table:      table,
		Visual:     visual,
	}
	return &out
}

func (r *ModelResourceModel) ToGetSDKType() *shared.ModelCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *ModelResourceModel) ToUpdateSDKType() *shared.ModelUpdate {
	var custom *shared.ModelUpdateCustom
	if r.Custom != nil {
		var query interface{}
		_ = json.Unmarshal([]byte(r.Custom.Query.ValueString()), &query)
		custom = &shared.ModelUpdateCustom{
			Query: query,
		}
	}
	var dbt *shared.ModelUpdateDbt
	if r.Dbt != nil {
		modelID := r.Dbt.ModelID.ValueString()
		dbt = &shared.ModelUpdateDbt{
			ModelID: modelID,
		}
	}
	isSchema := new(bool)
	if !r.IsSchema.IsUnknown() && !r.IsSchema.IsNull() {
		*isSchema = r.IsSchema.ValueBool()
	} else {
		isSchema = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	primaryKey := new(string)
	if !r.PrimaryKey.IsUnknown() && !r.PrimaryKey.IsNull() {
		*primaryKey = r.PrimaryKey.ValueString()
	} else {
		primaryKey = nil
	}
	var raw *shared.ModelUpdateRaw
	if r.Raw != nil {
		sql := r.Raw.SQL.ValueString()
		raw = &shared.ModelUpdateRaw{
			SQL: sql,
		}
	}
	var table *shared.ModelUpdateTable
	if r.Table != nil {
		name1 := r.Table.Name.ValueString()
		table = &shared.ModelUpdateTable{
			Name: name1,
		}
	}
	var visual *shared.ModelUpdateVisual
	if r.Visual != nil {
		var filter interface{}
		_ = json.Unmarshal([]byte(r.Visual.Filter.ValueString()), &filter)
		parentID := r.Visual.ParentID.ValueString()
		primaryLabel := r.Visual.PrimaryLabel.ValueString()
		secondaryLabel := r.Visual.SecondaryLabel.ValueString()
		visual = &shared.ModelUpdateVisual{
			Filter:         filter,
			ParentID:       parentID,
			PrimaryLabel:   primaryLabel,
			SecondaryLabel: secondaryLabel,
		}
	}
	out := shared.ModelUpdate{
		Custom:     custom,
		Dbt:        dbt,
		IsSchema:   isSchema,
		Name:       name,
		PrimaryKey: primaryKey,
		Raw:        raw,
		Table:      table,
		Visual:     visual,
	}
	return &out
}

func (r *ModelResourceModel) RefreshFromGetResponse(resp *shared.Model) {
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
